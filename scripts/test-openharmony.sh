#!/usr/bin/env bash
# Copyright (c) Tailscale Inc & contributors
# SPDX-License-Identifier: BSD-3-Clause
#
# Build and regression-test the OpenHarmony target with the pinned
# Xinlong-Wu/go-ohos toolchain. The compile phase is intentionally broad (it
# catches package and build-tag regressions); the runtime phase executes a
# focused set of tests and a tailscaled/CLI smoke test in DockerHarmony.

set -Eeuo pipefail

usage() {
  cat >&2 <<'EOF'
usage: scripts/test-openharmony.sh [all|compile|runtime]

all (the default) runs the complete OpenHarmony compile audit and the
DockerHarmony runtime checks. compile skips DockerHarmony, while runtime still
builds the binaries it needs before running them.

Useful overrides:
  GO_OHOS_ROOT       use an existing go-ohos GOROOT
  GO_OHOS_HOST_ARCH  select the host-tool asset (arm64 needs QEMU on amd64)
  GO_OHOS_QEMU       path to qemu-aarch64 for an ARM64 host-tool asset
  GO_OHOS_CACHE_DIR  cache directory for the toolchain and build cache
  OPENHARMONY_WORK_DIR  keep build artifacts in this directory
  OPENHARMONY_KEEP_WORK=1  retain an automatically-created work directory
  OPENHARMONY_DOCKER_PULL=always  refresh the pinned DockerHarmony image
EOF
  exit 2
}

mode="${1:-all}"
case "$mode" in
  all|compile|runtime) ;;
  -h|--help) usage ;;
  *) usage ;;
esac

script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd -P)"
repo_root="$(cd -- "$script_dir/.." && pwd -P)"
cd "$repo_root"

config_file="$repo_root/tool/openharmony/toolchain.env"
[[ -r "$config_file" ]] || { echo "OpenHarmony: missing $config_file" >&2; exit 1; }
# shellcheck disable=SC1090 # The path is fixed relative to the repository.
source "$config_file"

log() {
  printf 'openharmony: %s\n' "$*"
}

die() {
  printf 'openharmony: ERROR: %s\n' "$*" >&2
  exit 1
}

work_dir="${OPENHARMONY_WORK_DIR:-}"
created_work_dir=false
if [[ -z "$work_dir" ]]; then
  work_base="${OPENHARMONY_TMP_BASE:-${RUNNER_TEMP:-/var/tmp}}"
  if [[ ! -d "$work_base" || ! -w "$work_base" ]]; then
    work_base="${TMPDIR:-/tmp}"
  fi
  work_dir="$(mktemp -d "$work_base/tailscale-openharmony.XXXXXX")"
  created_work_dir=true
else
  mkdir -p "$work_dir"
  work_dir="$(cd -- "$work_dir" && pwd -P)"
fi

mkdir -p "$work_dir/bin" "$work_dir/tmp" "$work_dir/gotmp" "$work_dir/container-tmp"

daemon_container=""
cleanup() {
  status=$?
  trap - EXIT
  set +e
  if [[ -n "$daemon_container" ]]; then
    docker rm -f "$daemon_container" >/dev/null 2>&1
  fi
  if [[ "$created_work_dir" == true && "$status" == 0 && "${OPENHARMONY_KEEP_WORK:-}" != 1 ]]; then
    rm -rf -- "$work_dir"
  else
    log "work directory retained at $work_dir"
  fi
  exit "$status"
}
trap cleanup EXIT

# Keep module and build caches separate from the repository's Tailscale Go
# cache. CI maps these directories to actions/cache; local runs retain them
# under the go-ohos cache by default.
ohos_cache_dir="${GO_OHOS_CACHE_DIR:-${XDG_CACHE_HOME:-$HOME/.cache}/go-ohos}"
export GO_OHOS_CACHE_DIR="$ohos_cache_dir"
export GOMODCACHE="${GOMODCACHE:-$ohos_cache_dir/gomodcache}"
export GOCACHE="${GOCACHE:-$ohos_cache_dir/build-cache}"
export GOTMPDIR="$work_dir/gotmp"
export TMPDIR="$work_dir/tmp"
export GOOS=openharmony
export GOARCH=arm64
export CGO_ENABLED=0
export GOTOOLCHAIN=local

go_cmd="$repo_root/tool/go-ohos"
[[ -x "$go_cmd" ]] || die "$go_cmd is not executable"

log "toolchain: $($go_cmd version)"
log "target: GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=$CGO_ENABLED GOTOOLCHAIN=$GOTOOLCHAIN"
log "GOMODCACHE=$GOMODCACHE"
log "GOCACHE=$GOCACHE"

check_build_constraints() {
  local file expr
  while IFS= read -r -d '' file; do
    expr="$(sed -n 's|^[[:space:]]*//[[:space:]]*go:build[[:space:]]*||p' "$file" | head -n 1)"
    if [[ "$expr" == *'!openharmony'* ]]; then
      die "${file#"$repo_root/"} has a !openharmony build tag despite its _openharmony.go filename"
    fi
  done < <(find "$repo_root" -type f -name '*_openharmony.go' -print0)
  log "build-constraint audit passed"
}

build_production_binaries() {
  local name package
  declare -a names=(tailscale tailscaled derper k8s-operator)
  for name in "${names[@]}"; do
    package="./cmd/$name"
    log "building $package"
    "$go_cmd" build -trimpath -o "$work_dir/bin/$name" "$package"
  done
}

build_runtime_tests() {
  # Keep these tests focused on code that is meaningful in DockerHarmony's
  # mini rootfs. Tests that deliberately invoke a host 'go' command or require
  # a real TPM/TUN device are covered by the compile audit or are skipped by
  # their own capability checks.
  declare -a test_names=(version netmon netns tpm ipnlocal)
  declare -a test_packages=(
    ./version
    ./net/netmon
    ./net/netns
    ./feature/tpm
    ./ipn/ipnlocal
  )
  declare -a test_patterns=(
    'TestShortAllocs|TestIsValidLongWithTwoRepos|TestPrepExeNameForCmp|TestParse|TestAtLeast'
    'TestMonitorStartClose|TestMonitorJustClose|TestMonitorInjectEvent|TestMonitorInjectEventOnBus|TestGetState|TestInterfaceDiff|TestForeachInterface|TestStateString|TestEqual|TestPrefixesEqual'
    'TestIsLocalhost'
    'TestPropToString|TestMigrateStateToTPM'
    'TestViaTargetAllowed|TestFlagExpiredPeers|TestNextPeerExpiry|TestIsNotableNotify|TestMergePeerChangedPatch|TestExpandProxyArg'
  )
  local i name
  for i in "${!test_names[@]}"; do
    name="${test_names[$i]}"
    log "compiling runtime test $name (${test_packages[$i]})"
    "$go_cmd" test -c -trimpath -o "$work_dir/bin/$name.test" "${test_packages[$i]}"
  done
  RUNTIME_TEST_NAMES=("${test_names[@]}")
  RUNTIME_TEST_PATTERNS=("${test_patterns[@]}")
}

compile_phase() {
  check_build_constraints

  # Compile the non-OpenHarmony fallback files with the same go-ohos compiler.
  # This catches the class of filename/build-tag regressions that only appears
  # when a newly recognized GOOS suffix is added to the toolchain.
  log "checking non-OpenHarmony fallback packages"
  env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GOTOOLCHAIN=local \
    "$go_cmd" test -exec=true -run '^$' -count=1 ./version ./feature/tpm
  env GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 GOTOOLCHAIN=local \
    "$go_cmd" test -exec=true -run '^$' -count=1 ./version ./feature/tpm

  log "compiling all OpenHarmony packages"
  "$go_cmd" test -exec=true -run '^$' -count=1 ./...
  build_production_binaries
}

docker_image="${OPENHARMONY_DOCKER_IMAGE:?OPENHARMONY_DOCKER_IMAGE is not set in toolchain.env}"
docker_arch=""
docker_qemu=""
declare -a docker_mount_args

setup_docker() {
  command -v docker >/dev/null 2>&1 || die "Docker is required for runtime checks"
  local pull_policy="${OPENHARMONY_DOCKER_PULL:-if-missing}"
  case "$pull_policy" in
    always)
      log "pulling pinned DockerHarmony image"
      docker pull --platform linux/arm64 "$docker_image" >/dev/null
      ;;
    if-missing)
      if ! docker image inspect "$docker_image" >/dev/null 2>&1; then
        log "pulling pinned DockerHarmony image"
        docker pull --platform linux/arm64 "$docker_image" >/dev/null
      fi
      ;;
    *) die "OPENHARMONY_DOCKER_PULL must be always or if-missing" ;;
  esac

  local image_arch
  image_arch="$(docker image inspect --format '{{.Architecture}}' "$docker_image")"
  [[ "$image_arch" == arm64 ]] || die "DockerHarmony image architecture is $image_arch, expected arm64"
  docker_arch="$(docker info --format '{{.Architecture}}')"
  case "$docker_arch" in
    arm64|aarch64) docker_arch=arm64 ;;
    amd64|x86_64) docker_arch=amd64 ;;
    *) die "unsupported Docker daemon architecture: $docker_arch" ;;
  esac

  docker_mount_args=(
    --platform linux/arm64
    --volume "$work_dir:/work:rw"
    --volume "$work_dir/container-tmp:/data/local/tmp:rw"
    --env TMPDIR=/data/local/tmp
  )
  if [[ "$docker_arch" == amd64 ]]; then
    docker_qemu="${GO_OHOS_QEMU:-}"
    if [[ -z "$docker_qemu" ]]; then
      for candidate in qemu-aarch64-static qemu-aarch64; do
        if command -v "$candidate" >/dev/null 2>&1; then
          docker_qemu="$(command -v "$candidate")"
          break
        fi
      done
    fi
    [[ -n "$docker_qemu" && -x "$docker_qemu" ]] || die "DockerHarmony is arm64 but the Docker daemon is amd64; install qemu-aarch64-static or set GO_OHOS_QEMU"
    if command -v readlink >/dev/null 2>&1; then
      docker_qemu="$(readlink -f "$docker_qemu")"
    fi
    docker_mount_args+=(--volume "$docker_qemu:/qemu-aarch64:ro")
    log "DockerHarmony execution: amd64 host with $docker_qemu"
  else
    log "DockerHarmony execution: native arm64"
  fi
}

docker_run_target() {
  local program="$1"
  shift
  if [[ "$docker_arch" == arm64 ]]; then
    docker run --rm "${docker_mount_args[@]}" "$docker_image" "$program" "$@"
  else
    docker run --rm "${docker_mount_args[@]}" --entrypoint /qemu-aarch64 \
      "$docker_image" "$program" "$program" "$@"
  fi
}

docker_exec_target() {
  local container="$1" program="$2"
  shift 2
  if [[ "$docker_arch" == arm64 ]]; then
    docker exec "$container" "$program" "$@"
  else
    docker exec "$container" /qemu-aarch64 "$program" "$program" "$@"
  fi
}

run_runtime_tests() {
  local i name pattern
  for i in "${!RUNTIME_TEST_NAMES[@]}"; do
    name="${RUNTIME_TEST_NAMES[$i]}"
    pattern="${RUNTIME_TEST_PATTERNS[$i]}"
    log "running $name tests in DockerHarmony"
    docker_run_target "/work/bin/$name.test" -test.v -test.run "$pattern"
  done
}

run_daemon_smoke() {
  local socket=/data/local/tmp/tailscaled.sock
  local status_file="$work_dir/container-tmp/status.json"
  local status_err="$work_dir/container-tmp/status.err"
  local daemon_log="$work_dir/container-tmp/daemon.log"
  daemon_container="tailscale-oh-openharmony-${$}-${RANDOM}"

  log "starting tailscaled userspace-networking smoke test in DockerHarmony"
  if [[ "$docker_arch" == arm64 ]]; then
    docker run -d --name "$daemon_container" "${docker_mount_args[@]}" "$docker_image" \
      /work/bin/tailscaled \
      --tun=userspace-networking --state=mem: --statedir=/data/local/tmp/state \
      --socket="$socket" --no-logs-no-support >/dev/null
  else
    docker run -d --name "$daemon_container" "${docker_mount_args[@]}" --entrypoint /qemu-aarch64 "$docker_image" \
      /work/bin/tailscaled /work/bin/tailscaled \
      --tun=userspace-networking --state=mem: --statedir=/data/local/tmp/state \
      --socket="$socket" --no-logs-no-support >/dev/null
  fi
  local attempt
  for attempt in $(seq 1 60); do
    if ! docker inspect -f '{{.State.Running}}' "$daemon_container" 2>/dev/null | grep -qx true; then
      docker logs "$daemon_container" > "$daemon_log" 2>&1 || true
      die "tailscaled exited before its local API became available; see $daemon_log"
    fi
    if docker_exec_target "$daemon_container" /work/bin/tailscale \
      --socket="$socket" status --json >"$status_file" 2>"$status_err"; then
      break
    fi
    if [[ "$attempt" == 60 ]]; then
      docker logs "$daemon_container" > "$daemon_log" 2>&1 || true
      die "timed out waiting for tailscaled local API; see $daemon_log"
    fi
    sleep 1
  done

  if command -v jq >/dev/null 2>&1; then
    [[ "$(jq -r '.BackendState' "$status_file")" == NeedsLogin ]] || die "unexpected daemon BackendState; see $status_file"
    [[ "$(jq -r '.Self.OS' "$status_file")" == OpenHarmony ]] || die "unexpected daemon Self.OS; see $status_file"
  else
    grep -q '"BackendState"[[:space:]]*:[[:space:]]*"NeedsLogin"' "$status_file" || die "unexpected daemon BackendState; see $status_file"
    grep -q '"OS"[[:space:]]*:[[:space:]]*"OpenHarmony"' "$status_file" || die "unexpected daemon Self.OS; see $status_file"
  fi
  log "daemon reported BackendState=NeedsLogin and Self.OS=OpenHarmony"

  docker logs "$daemon_container" > "$daemon_log" 2>&1 || true
  docker stop --time 10 "$daemon_container" >/dev/null
  local exit_code
  exit_code="$(docker inspect -f '{{.State.ExitCode}}' "$daemon_container")"
  [[ "$exit_code" == 0 ]] || die "tailscaled exited with status $exit_code; see $daemon_log"
  docker rm "$daemon_container" >/dev/null
  daemon_container=""
  log "daemon shut down cleanly"
  log "DockerHarmony limitations: no TPM/TUN device and shared host kernel are expected; this smoke test covers the OpenHarmony rootfs and userspace path"
}

runtime_phase() {
  build_production_binaries
  build_runtime_tests
  setup_docker
  # Verify that the selected Docker/QEMU combination can execute the target
  # rootfs before spending time on the package tests.
  docker_run_target /bin/true
  run_runtime_tests
  run_daemon_smoke
}

case "$mode" in
  compile)
    compile_phase
    ;;
  runtime)
    runtime_phase
    ;;
  all)
    compile_phase
    build_runtime_tests
    setup_docker
    docker_run_target /bin/true
    run_runtime_tests
    run_daemon_smoke
    ;;
esac

log "$mode phase completed successfully"
