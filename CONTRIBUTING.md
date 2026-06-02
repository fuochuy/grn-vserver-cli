# Contributing to GreenNode CLI

Thank you for your interest in contributing! This guide covers everything you need to get started.

## Table of contents

- [Code of conduct](#code-of-conduct)
- [Getting started](#getting-started)
- [Development workflow](#development-workflow)
- [Making changes](#making-changes)
- [Testing](#testing)
- [Changelog fragments](#changelog-fragments)
- [Pull request checklist](#pull-request-checklist)
- [Release process](#release-process)

---

## Code of conduct

Be respectful and constructive. We follow the [Contributor Covenant](https://www.contributor-covenant.org/version/2/1/code_of_conduct/).

---

## Getting started

### Prerequisites

- Go 1.22 or later — [download](https://go.dev/dl/)
- Git

### Fork and clone

```bash
git clone https://github.com/<your-username>/greennode-cli.git
cd greennode-cli/go
```

### Build

```bash
# Build for the current platform
make build           # → ./grn

# Cross-compile all release targets
make build-all
```

### Verify

```bash
./grn --version
./grn --help
```

---

## Development workflow

```
main
 └── feature/<short-description>   ← your branch
 └── fix/<issue-number>-<desc>
 └── docs/<what-you-are-documenting>
```

1. Branch from `main`.
2. Make focused, atomic commits.
3. Open a pull request against `main`.

---

## Making changes

### Project layout

```
go/
├── cmd/
│   ├── root.go             # global flags, version
│   ├── configure/          # grn configure ...
│   ├── mcp/                # grn mcp
│   └── vserver/
│       ├── vserver.go      # grn vserver (top-level)
│       ├── server/         # grn vserver server ...
│       ├── volume/         # grn vserver volume ...
│       ├── vpc/            # grn vserver vpc ...
│       ├── subnet/         # grn vserver subnet ...
│       ├── secgroup/       # grn vserver secgroup ...
│       ├── flavor/         # grn vserver flavor ...
│       ├── image/          # grn vserver image ...
│       └── volumetype/     # grn vserver volume-type ...
└── internal/
    ├── auth/               # OAuth2 token handling
    ├── client/             # base HTTP client
    ├── config/             # credential file read/write
    ├── formatter/          # json/table/text output
    ├── validator/          # ID format validation
    └── vserverclient/      # vServer API client + zone helpers
```

### Adding a new command

1. Create a package under `cmd/vserver/<resource>/`.
2. Define the `cobra.Command` with `Use`, `Short`, and `RunE`.
3. Register it in the parent `<resource>.go` file via `AddCommand`.
4. Add shell completion in `zcompletion.go` if the resource has IDs.
5. Write the corresponding doc page in `docs/commands/vserver/`.

### Coding conventions

- `RunE` for all commands (never `Run`) so errors propagate correctly.
- Use `cmd.Flags().MarkFlagRequired(...)` for mandatory flags; panic with `BUG:` prefix if it fails (signals a programming error, not a runtime error).
- Output goes through `outputResult(cmd, cfg, result)` — respects `--output` and `--query`.
- IDs are validated with `validator.ValidateID(id, flagName)` before any API call.
- Destructive commands (`delete`) print a preview and prompt before acting unless `--force` is set.

---

## Testing

```bash
# Run all tests
make test

# Vet
make vet

# Format (required before committing)
make fmt
```

There are currently no automated integration tests; manual testing against a VNG Cloud project is the primary validation path.

---

## Changelog fragments

Every pull request that changes user-visible behaviour must include a changelog entry.

Add a file to `.changes/` named `<type>-<short-description>.md`:

| Type | When to use |
|------|-------------|
| `feature` | New command or flag |
| `fix` | Bug fix |
| `change` | Modified default behaviour or output format |
| `docs` | Documentation-only change |
| `internal` | Refactor, dependency update, CI — no user impact |

**File format** (`.changes/feature-subnet-commands.md`):

```markdown
---
type: feature
---

Added `grn vserver subnet` commands: `list`, `get`, `create`, `delete`.
```

These fragments are merged into `CHANGELOG.md` during the release process.

---

## Pull request checklist

Before requesting a review:

- [ ] `make fmt` passes without diff
- [ ] `make vet` passes
- [ ] `make test` passes
- [ ] `make build` produces a working binary
- [ ] New commands have a corresponding doc page in `docs/commands/vserver/`
- [ ] A changelog fragment exists in `.changes/` (skip for `docs`-only PRs)
- [ ] Commit messages are clear and reference an issue if applicable

---

## Release process

Releases are cut by the maintainers using the `scripts/bump-version` script:

```bash
# Bump patch version (e.g. 1.3.2 → 1.3.3)
./scripts/bump-version patch

# Bump minor version (e.g. 1.3.2 → 1.4.0)
./scripts/bump-version minor

# Bump major version (e.g. 1.3.2 → 2.0.0)
./scripts/bump-version major
```

The script:
1. Updates `cliVersion` in `go/cmd/root.go`.
2. Merges `.changes/` fragments into `CHANGELOG.md`.
3. Creates a git commit `release: vX.Y.Z` and a matching tag.

After that, push the commit and tag:

```bash
git push && git push --tags
```

CI picks up the tag and publishes the release binaries automatically.
