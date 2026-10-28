# Upstream synchronization conflict report

> [!CAUTION]
> This is a temporary staging file. Do not merge this PR until the signed upstream release has been merged, every conflict has been resolved, this file has been removed, and the release manifest has been updated.

- Upstream release: [v1.104.4](https://github.com/tailscale/tailscale/releases/tag/v1.104.4)
- Upstream commit: `27047635f237d1f3e0b08c0f487674757f30d2f7`
- Verified signing-key fingerprint: `SHA256:4X7QKZpmXLkns4z0XcXcUmifM/jruCmFyJnJIVmdUNo`
- Base branch: `main`
- Base commit: `40a8ac55b87a10070ea76d0cf298449db67b8b63`
- Conflicting files: `12`
- Go/toolchain metadata changed: `false`

## Manual resolution

1. Check out this PR branch.
2. Fetch the verified tag and merge its peeled commit without auto-committing:

   ```sh
   git fetch https://github.com/tailscale/tailscale.git refs/tags/v1.104.4:refs/tags/v1.104.4
   git merge --no-ff --no-commit 27047635f237d1f3e0b08c0f487674757f30d2f7
   ```

3. Resolve every conflict while preserving the fork-specific OpenHarmony adaptations. Do not apply a blanket `ours` or `theirs` resolution.
4. Remove this report as part of the signed-off upstream merge commit.
5. Update `.github/upstream-release.json`, run the required validation, and add the signed-off release-preparation commit.
6. Mark the PR ready only after the branch is complete and reviewed.

## Conflicting files

```text
cmd/tailscaled/depaware.txt
flake.nix
flakehashes.json
ipn/ipnlocal/local.go
net/netmon/netmon_windows.go
net/netns/netns_windows_test.go
shell.nix
ssh/tailssh/c2n.go
ssh/tailssh/incubator_test.go
ssh/tailssh/listen.go
ssh/tailssh/localstate_test.go
tsnet/listenssh_test.go
```
