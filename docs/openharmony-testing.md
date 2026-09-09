# OpenHarmony build and tests

The OpenHarmony lane uses the reviewed binary distribution from
[Xinlong-Wu/go-ohos](https://github.com/Xinlong-Wu/go-ohos), not the upstream
Tailscale Go toolchain. The pinned release, archive checksums, reviewed
`go.mod` version, and reviewed `go.toolchain.rev` are in
`tool/openharmony/toolchain.env`.

Run the complete local lane with:

```sh
./scripts/test-openharmony.sh all
```

The script also accepts `compile` (compile audit only) and `runtime` (build the
runtime inputs and execute focused tests in DockerHarmony). An existing local
toolchain can be selected with `GO_OHOS_ROOT`, for example:

```sh
GO_OHOS_ROOT=/root/go_ohos/amd64/go1.26.8-ohos.1 \
  ./scripts/test-openharmony.sh all
```

When an ARM64 host-tool release is being used on an x86-64 machine, provide a
user-mode QEMU binary and identify the selected host asset:

```sh
GO_OHOS_ROOT=/root/go_ohos/arm64/go1.26.8-ohos.1 \
GO_OHOS_HOST_ARCH=arm64 \
GO_OHOS_QEMU=/path/to/qemu-aarch64 \
  ./scripts/test-openharmony.sh all
```

The compile phase builds every package for `openharmony/arm64`, checks the
non-OpenHarmony fallback packages, and builds the main command binaries. The
runtime phase runs representative version, networking, namespace, TPM-store,
and `ipnlocal` tests plus a `tailscaled` userspace-networking/CLI smoke test in
the pinned `ghcr.io/hqzing/dockerharmony:6.1` image.

DockerHarmony is a mini rootfs that shares the host kernel. It does not provide
the complete OpenHarmony framework, a hardware TPM, or a kernel TUN device;
those capabilities require a real OpenHarmony device or emulator and are not
treated as failures of this lane. The smoke test uses static `CGO_ENABLED=0`
binaries, so it validates the rootfs and userspace behavior rather than dynamic
musl linkage.
