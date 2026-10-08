#!/usr/bin/env bash
# Copyright (c) Tailscale Inc & contributors
# SPDX-License-Identifier: BSD-3-Clause
#
# Prepare a clean merge of the latest formal tailscale/tailscale GitHub
# Release. This script creates local commits only. The calling workflow is
# responsible for pushing the branch and opening the pull request.

set -Eeuo pipefail

script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd -P)"
repo_root="$(cd -- "$script_dir/.." && pwd -P)"
cd "$repo_root"

upstream_repository="${UPSTREAM_REPOSITORY:-tailscale/tailscale}"
downstream_repository="${DOWNSTREAM_REPOSITORY:-${GITHUB_REPOSITORY:-Xinlong-Wu/tailscale-oh}}"
base_branch="${BASE_BRANCH:-main}"
sync_date="${SYNC_DATE:-$(date -u +%F)}"
sync_branch="sync/upstream-${sync_date}"
upstream_url="https://github.com/${upstream_repository}.git"
trusted_signers_file="$repo_root/.github/upstream-release-signers"

output_file="${GITHUB_OUTPUT:-/dev/null}"
summary_file="${GITHUB_STEP_SUMMARY:-/dev/null}"
report_file="${SYNC_REPORT_FILE:-${TMPDIR:-/tmp}/upstream-sync-report.md}"
issue_title_file="${SYNC_ISSUE_TITLE_FILE:-${TMPDIR:-/tmp}/upstream-sync-issue-title.txt}"

write_output() {
	printf '%s=%s\n' "$1" "$2" >> "$output_file"
}

write_summary() {
	if [[ "$summary_file" != /dev/null ]]; then
		cat >> "$summary_file"
	else
		cat >/dev/null
	fi
}

block() {
	local title="$1" body_file
	body_file="$(mktemp)"
	cat > "$body_file"
	printf '%s\n' "$title" > "$issue_title_file"
	{
		printf '## Automated upstream synchronization stopped\n\n'
		cat "$body_file"
		if [[ -n "${GITHUB_SERVER_URL:-}" && -n "${GITHUB_REPOSITORY:-}" && -n "${GITHUB_RUN_ID:-}" ]]; then
			printf '\nWorkflow run: %s/%s/actions/runs/%s\n' \
				"$GITHUB_SERVER_URL" "$GITHUB_REPOSITORY" "$GITHUB_RUN_ID"
		fi
	} > "$report_file"
	rm -f -- "$body_file"
	cat "$report_file" >&2
	write_summary < "$report_file"
	write_output action blocked
	exit 1
}

noop() {
	local message="$1"
	write_output action noop
	{
		printf '## Automated upstream synchronization\n\n'
		printf '%s\n' "$message"
	} | write_summary
	printf '%s\n' "$message"
	exit 0
}

if [[ "$upstream_repository" != "tailscale/tailscale" ]]; then
	block "Automated upstream sync has an invalid repository" <<EOF
The configured upstream repository is \`$upstream_repository\`, but only
\`tailscale/tailscale\` is permitted.
EOF
fi

if [[ "$downstream_repository" != "Xinlong-Wu/tailscale-oh" ]]; then
	block "Automated upstream sync has an invalid downstream repository" <<EOF
The workflow is running for \`$downstream_repository\`. It is intentionally
restricted to \`Xinlong-Wu/tailscale-oh\`.
EOF
fi

if [[ -z "${SYNC_PUSH_TOKEN:-}" ]]; then
	block "Automated upstream sync needs repository configuration" <<'EOF'
The `UPSTREAM_SYNC_TOKEN` Actions secret is missing. Configure a fine-grained
token with read/write access to Contents and Pull requests for this repository.
A token distinct from `GITHUB_TOKEN` is required so the created pull request
triggers the normal pull-request CI workflows.
EOF
fi

command -v gh >/dev/null 2>&1 || block "Automated upstream sync runner is missing gh" <<'EOF'
The GitHub CLI (`gh`) is required to discover the latest formal upstream
release.
EOF
command -v jq >/dev/null 2>&1 || block "Automated upstream sync runner is missing jq" <<'EOF'
`jq` is required to read and update the upstream release manifest.
EOF

release_fields="$(gh api "repos/${upstream_repository}/releases/latest" \
	--jq '[.tag_name, (.draft | tostring), (.prerelease | tostring), .html_url, .published_at] | @tsv')"
IFS=$'\t' read -r release_tag release_draft release_prerelease release_url release_published_at <<<"$release_fields"

if [[ ! "$release_tag" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
	block "Automated upstream sync found an invalid release tag" <<EOF
The latest GitHub Release returned the unexpected tag \`$release_tag\`.
Only stable tags in the form \`vMAJOR.MINOR.PATCH\` are accepted.
EOF
fi
if [[ "$release_draft" != false || "$release_prerelease" != false ]]; then
	block "Automated upstream sync found a non-stable release" <<EOF
The latest release endpoint returned draft=$release_draft and
prerelease=$release_prerelease for \`$release_tag\`. No branch was pushed.
EOF
fi

current_tag="$(jq -er '.upstream_tag | select(type == "string" and length > 0)' .github/upstream-release.json)"
if [[ "$current_tag" == "$release_tag" ]]; then
	noop "The latest upstream release, \`$release_tag\`, is already recorded on \`$base_branch\`."
fi

newest_version="$(printf '%s\n' "${current_tag#v}" "${release_tag#v}" | sort -V | tail -n 1)"
if [[ "$newest_version" != "${release_tag#v}" ]]; then
	block "Automated upstream sync detected a release regression" <<EOF
The fork records \`$current_tag\`, but GitHub reports \`$release_tag\` as the
latest release. Refusing to synchronize an older version.
EOF
fi

open_sync_pr="$(gh pr list \
	--repo "$downstream_repository" \
	--base "$base_branch" \
	--state open \
	--limit 100 \
	--json number,headRefName,title,url \
	--jq '.[] | select(.headRefName | startswith("sync/upstream-")) | [.number, .headRefName, .title, .url] | @tsv' | sed -n '1p')"
if [[ -n "$open_sync_pr" ]]; then
	IFS=$'\t' read -r open_pr_number _open_pr_branch open_pr_title open_pr_url <<<"$open_sync_pr"
	noop "A synchronization PR is already open: #${open_pr_number} (${open_pr_title}) — ${open_pr_url}. Resolve it before preparing \`$release_tag\`."
fi

if [[ -n "$(git status --porcelain)" ]]; then
	block "Automated upstream sync found a dirty checkout" <<'EOF'
The checked-out repository contains uncommitted changes. No synchronization
branch was prepared.
EOF
fi

git fetch --no-tags --force origin \
	"refs/heads/${base_branch}:refs/remotes/origin/${base_branch}"
base_ref="refs/remotes/origin/${base_branch}"

if git ls-remote --exit-code origin "refs/heads/${sync_branch}" >/dev/null 2>&1; then
	block "Automated upstream sync found an existing branch" <<EOF
The remote branch \`$sync_branch\` already exists without an open sync PR.
It was not overwritten. Create a PR from it or remove it after reviewing its
contents, then rerun the workflow.
EOF
fi

git fetch --no-tags --force "$upstream_url" \
	"refs/tags/${release_tag}:refs/tags/${release_tag}"

if [[ ! -r "$trusted_signers_file" ]]; then
	block "Automated upstream sync is missing its signer trust list" <<EOF
The trusted release-signing fingerprint file is not readable:
\`$trusted_signers_file\`.
EOF
fi

if [[ "$(git cat-file -t "refs/tags/${release_tag}")" != tag ]]; then
	block "Automated upstream sync found an unsigned release tag" <<EOF
\`$release_tag\` is not an annotated tag. Formal upstream releases must use a
signed annotated tag.
EOF
fi

verify_output="$(git -c gpg.ssh.allowedSignersFile=/dev/null \
	verify-tag --raw "$release_tag" 2>&1 || true)"
release_fingerprint="$(awk '/^Good "git" signature with ED25519 key SHA256:/ {print $NF; exit}' <<<"$verify_output")"
if [[ -z "$release_fingerprint" ]] || \
	! grep -Fqx -- "$release_fingerprint" "$trusted_signers_file"; then
	block "Automated upstream sync found an untrusted release signature" <<EOF
The signature on \`$release_tag\` did not match a fingerprint in
\`.github/upstream-release-signers\`.

Observed verification output:

\`\`\`
$verify_output
\`\`\`

Review any signing-key rotation manually before changing the trust list.
EOF
fi

release_commit="$(git rev-parse "${release_tag}^{commit}")"

base_go_version="$(git show "${base_ref}:go.mod" | awk '$1 == "go" {print $2; exit}')"
release_go_version="$(git show "${release_commit}:go.mod" | awk '$1 == "go" {print $2; exit}')"
base_toolchain_version="$(git show "${base_ref}:go.toolchain.version" | tr -d '[:space:]')"
release_toolchain_version="$(git show "${release_commit}:go.toolchain.version" | tr -d '[:space:]')"
base_toolchain_rev="$(git show "${base_ref}:go.toolchain.rev" | tr -d '[:space:]')"
release_toolchain_rev="$(git show "${release_commit}:go.toolchain.rev" | tr -d '[:space:]')"

toolchain_changed=false
if [[ "$base_go_version" != "$release_go_version" || \
	"$base_toolchain_version" != "$release_toolchain_version" || \
	"$base_toolchain_rev" != "$release_toolchain_rev" ]]; then
	toolchain_changed=true
fi

git switch --force-create "$sync_branch" "$base_ref"
git config user.name "github-actions[bot]"
git config user.email "41898282+github-actions[bot]@users.noreply.github.com"

set +e
# The annotated tag was verified above. Merge its peeled commit so Git does not
# attempt a second SSH trust check while constructing the merge commit.
git merge --no-ff --no-commit "$release_commit"
merge_status=$?
set -e
if (( merge_status != 0 )); then
	conflicting_files="$(git diff --name-only --diff-filter=U)"
	conflict_count="$(wc -l <<<"$conflicting_files" | tr -d '[:space:]')"
	conflicting_preview="$(sed -n '1,100p' <<<"$conflicting_files")"
	if (( conflict_count > 100 )); then
		conflicting_preview+=$'\n'
		conflicting_preview+="... and $((conflict_count - 100)) more conflicting files"
	fi
	git merge --abort || true
	block "Automated upstream sync has merge conflicts for ${release_tag}" <<EOF
The signed upstream release tag could not be merged cleanly into
\`$base_branch\`. No branch was pushed and no PR was created.

Conflicting files:

\`\`\`
$conflicting_preview
\`\`\`

Resolve this release manually so the fork-specific OpenHarmony adaptations are
preserved.
EOF
fi

if ! git rev-parse --verify --quiet MERGE_HEAD >/dev/null; then
	block "Automated upstream sync produced no merge commit for ${release_tag}" <<EOF
Git reported no merge commit after preparing \`$release_tag\`. The release
manifest still points to \`$current_tag\`, so this state requires manual review.
EOF
fi

git commit --signoff -m "Merge upstream release ${release_tag}"

version="$(tr -d '[:space:]' < VERSION.txt)"
if [[ "v${version}" != "$release_tag" ]]; then
	block "Automated upstream sync found a version mismatch for ${release_tag}" <<EOF
After the merge, \`VERSION.txt\` contains \`$version\` instead of
\`${release_tag#v}\`. No branch was pushed.
EOF
fi

module_path="$(awk '$1 == "module" {print $2; exit}' go.mod)"
if [[ "$module_path" != "github.com/Xinlong-Wu/tailscale-oh" ]]; then
	block "Automated upstream sync changed the fork module path for ${release_tag}" <<EOF
After the merge, the root module is \`$module_path\`. The fork requires
\`github.com/Xinlong-Wu/tailscale-oh\`. No global import rewrite was attempted;
this release requires manual conflict resolution.
EOF
fi

manifest_tmp=".github/upstream-release.json.tmp"
jq -n \
	--arg upstream_repository "$upstream_repository" \
	--arg upstream_tag "$release_tag" \
	--arg upstream_commit "$release_commit" \
	'{
	  upstream_repository: $upstream_repository,
	  upstream_tag: $upstream_tag,
	  upstream_commit: $upstream_commit,
	  prerelease: false
	}' > "$manifest_tmp"
chmod 0644 "$manifest_tmp"
mv -- "$manifest_tmp" .github/upstream-release.json

git add .github/upstream-release.json
git commit --signoff -m "sync: prepare ${release_tag} release"

if ! git diff --check "$base_ref"...HEAD; then
	block "Automated upstream sync found whitespace errors for ${release_tag}" <<EOF
The prepared synchronization branch fails \`git diff --check\`. No branch was
pushed. Review the merge result manually.
EOF
fi

write_output action prepared
write_output branch "$sync_branch"
write_output release_tag "$release_tag"
write_output release_url "$release_url"
write_output release_published_at "$release_published_at"
write_output upstream_commit "$release_commit"
write_output release_fingerprint "$release_fingerprint"
write_output toolchain_changed "$toolchain_changed"
write_output base_go_version "$base_go_version"
write_output release_go_version "$release_go_version"
write_output base_toolchain_version "$base_toolchain_version"
write_output release_toolchain_version "$release_toolchain_version"
write_output base_toolchain_rev "$base_toolchain_rev"
write_output release_toolchain_rev "$release_toolchain_rev"

{
	printf '## Prepared upstream synchronization\n\n'
	printf -- '- Release: [%s](%s)\n' "$release_tag" "$release_url"
	# shellcheck disable=SC2016 # Backticks are intentional Markdown delimiters.
	printf -- '- Upstream commit: `%s`\n' "$release_commit"
	# shellcheck disable=SC2016 # Backticks are intentional Markdown delimiters.
	printf -- '- Branch: `%s`\n' "$sync_branch"
	# shellcheck disable=SC2016 # Backticks are intentional Markdown delimiters.
	printf -- '- Toolchain metadata changed: `%s`\n' "$toolchain_changed"
} | write_summary
