# condition-gitea-actions

[![Latest Release](https://img.shields.io/github/v/release/SemRels/condition-gitea-actions?label=version&color=blue)](https://github.com/SemRels/condition-gitea-actions/releases/latest)
[![CI](https://github.com/SemRels/condition-gitea-actions/actions/workflows/ci.yml/badge.svg)](https://github.com/SemRels/condition-gitea-actions/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/SemRels/condition-gitea-actions)](LICENSE)
[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/SemRels/condition-gitea-actions/badge)](https://scorecard.dev/viewer/?uri=github.com/SemRels/condition-gitea-actions)
[![REUSE status](https://api.reuse.software/badge/github.com/SemRels/condition-gitea-actions)](https://api.reuse.software/info/github.com/SemRels/condition-gitea-actions)

Allows releases only when semrel is running inside Gitea Actions.

This plugin is distributed as the standalone Go binary `semrel-plugin-condition-gitea-actions`. Semrel executes the binary as a subprocess, provides plugin configuration through `SEMREL_PLUGIN_*` environment variables, provides release context through `SEMREL_*` environment variables, reads standard output, and treats exit code `0` as success and any non-zero exit code as failure. Install the binary in `~/.semrel/plugins/` or anywhere on your `$PATH`.

## Installation

### Binary

```bash
go install github.com/SemRels/condition-gitea-actions/cmd/plugin@latest
```

### Docker

Pre-built, multi-platform images (linux/amd64, linux/arm64) are published to the GitHub Container Registry on every release:

```bash
docker pull ghcr.io/semrels/condition-gitea-actions:latest
```

Images are signed with [cosign](https://github.com/sigstore/cosign) and include a full SBOM attestation. Verify the signature:

```bash
cosign verify ghcr.io/semrels/condition-gitea-actions:latest \
  --certificate-identity-regexp 'https://github.com/SemRels/condition-gitea-actions/.github/workflows/container.yml.*' \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com
```


## Configuration

```yaml
plugins:
  - uses: "@semrel/condition-gitea-actions"
    phase: condition
```

## Options

This plugin has no options. It relies on CI-provided environment state.

## `SEMREL_*` release context used

This plugin does not consume any `SEMREL_*` release context variables directly.

## Example behavior

The plugin checks the CI environment and succeeds when `GITEA_ACTIONS=true`. Outside Gitea Actions it exits non-zero to stop the release.

## License

Apache-2.0
