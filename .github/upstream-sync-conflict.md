# Upstream synchronization conflict report

> [!CAUTION]
> This is a temporary staging file. Do not merge this PR until the signed upstream release has been merged, every conflict has been resolved, this file has been removed, and the release manifest has been updated.

- Upstream release: [v1.104.2](https://github.com/tailscale/tailscale/releases/tag/v1.104.2)
- Upstream commit: `e26db7a5797683ec99ad1516c40b628c7f2dfd71`
- Verified signing-key fingerprint: `SHA256:4X7QKZpmXLkns4z0XcXcUmifM/jruCmFyJnJIVmdUNo`
- Base branch: `main`
- Base commit: `a4668fb55d66a5559286c7c4e7bfb7b87a9e3c82`
- Conflicting files: `3`
- Go/toolchain metadata changed: `true`

## Manual resolution

1. Check out this PR branch.
2. Fetch the verified tag and merge its peeled commit without auto-committing:

   ```sh
   git fetch https://github.com/tailscale/tailscale.git refs/tags/v1.104.2:refs/tags/v1.104.2
   git merge --no-ff --no-commit e26db7a5797683ec99ad1516c40b628c7f2dfd71
   ```

3. Resolve every conflict while preserving the fork-specific OpenHarmony adaptations. Do not apply a blanket `ours` or `theirs` resolution.
4. Remove this report as part of the signed-off upstream merge commit.
5. Update `.github/upstream-release.json`, run the required validation, and add the signed-off release-preparation commit.
6. Mark the PR ready only after the branch is complete and reviewed.

## Conflicting files

```text
flake.nix
flakehashes.json
shell.nix
```
