# Automated upstream release synchronization

The `Sync latest upstream release` workflow checks the latest formal GitHub
Release from `tailscale/tailscale` every third UTC day-of-month and can also be
run manually.

When a newer stable release exists, the workflow:

1. verifies that its annotated tag has an SSH signature from a fingerprint in
   `.github/upstream-release-signers`;
2. cleanly merges the signed tag into `sync/upstream-YYYY-MM-DD`, preserving
   the upstream and fork histories;
3. updates `.github/upstream-release.json`;
4. pushes only the synchronization branch; and
5. opens a pull request to `main` without enabling auto-merge.

The automation never pushes the fetched upstream tag to the fork and never
creates a fork tag or GitHub Release. Existing post-merge automation remains
responsible for publication after manual merge and successful same-commit
`main` CI.

## Required repository secret

Configure an Actions secret named `UPSTREAM_SYNC_TOKEN`. Use a fine-grained
personal access token restricted to `Xinlong-Wu/tailscale-oh` with:

- Contents: Read and write
- Pull requests: Read and write
- Workflows: Read and write, because an upstream release can change files
  under `.github/workflows/`

The workflow deliberately does not use `GITHUB_TOKEN` to push the branch or
open the PR. GitHub suppresses most workflow events created by `GITHUB_TOKEN`;
using a separate token ensures that the synchronization PR starts the normal
pull-request CI.

## Safety and expected manual work

- Drafts and prereleases are ignored by using GitHub's latest-release API and
  by validating the returned release metadata.
- An existing `sync/upstream-*` PR prevents another synchronization PR from
  being opened.
- Unknown signing keys, merge conflicts, dirty checkouts, version mismatches,
  and existing same-day branches stop before anything is pushed. The workflow
  creates or updates an issue describing the blocker.
- Go and Tailscale toolchain metadata changes do not prevent the PR from being
  opened. They are shown prominently in its body, and the existing OpenHarmony
  CI is expected to fail until a compatible `Xinlong-Wu/go-ohos` toolchain has
  been reviewed and configured. The scheduled workflow does not inspect or
  download a local toolchain and has no toolchain-compatibility gate.
- The merge is intentionally automatic only when Git reports no conflicts.
  CI and human review must still catch semantic conflicts, newly introduced
  upstream-only workflows, module-path changes, generated files, and
  OpenHarmony regressions.

GitHub cron schedules are based on days of the month rather than a persistent
72-hour timer, so `23 4 */3 * *` means every third UTC day-of-month. The
workflow is idempotent when the latest release is already synchronized.
