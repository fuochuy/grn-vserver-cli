# Changelog

All notable changes to `grn` CLI are documented here.

The format follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).
Versions follow [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

## [Unreleased]

---

## [1.3.2] - 2026-06-02

### Added
- `vserver subnet` commands: `list`, `get`, `create`, `delete`
- `vserver volume` commands: `list`, `get`, `create`, `delete`
- `vserver secgroup` commands: `list`, `get`, `create`, `delete`
- `vserver secgroup rule` commands: `list`, `get`, `create`, `delete`
- Shell completions for `secgroup`, `subnet`, `volume`, and `vpc` commands
- Full command documentation for all new resources

---

## [1.3.1] - 2025-10-15

### Added
- `vserver vpc` commands: `list`, `get`, `create`, `delete`
- `vserver flavor list-families` and `flavor list-codes` sub-commands
- JMESPath `--query` flag for all list and get commands
- `--color` flag (`on`, `off`, `auto`) for colored terminal output

### Changed
- `flavor list` now filters out sold-out flavors (`remainingVms <= 1`) automatically

---

## [1.3.0] - 2025-08-01

### Added
- `vserver image list` with `--type os|gpu` and `--image-version` filter
- `vserver volume-type list` with two-step zone → type lookup
- `vserver volume resize` — expand size and/or change volume type
- `vserver server resize` — change server flavor

### Fixed
- Pagination now correctly handles last-page boundary conditions

---

## [1.2.0] - 2025-05-10

### Added
- Built-in MCP server (`grn mcp`) exposing all commands as AI-assistant tools
- Auto-detect project ID from credentials during `grn configure`
- `--debug` global flag for verbose HTTP logging
- `--endpoint-url` global flag to override the API base URL

### Changed
- `grn configure` wizard now validates credentials on save

---

## [1.1.0] - 2025-02-20

### Added
- `vserver server` commands: `list`, `get`, `create`, `start`, `stop`, `reboot`, `delete`
- `vserver flavor list` with `--family` and `--code` filters
- `grn configure` wizard (Client ID / Client Secret / region / project ID)
- `grn configure list`, `get`, `set` for programmatic config management
- `--output` flag: `json` (default), `table`, `text`
- `--no-verify-ssl` and connection-timeout global flags

---

## [1.0.0] - 2024-11-01

### Added
- Initial release of `grn` CLI
- Authentication via VNG Cloud IAM service accounts (OAuth2 client credentials)
- `grn configure` for credential and region setup
- Cross-compiled binaries for Linux (amd64/arm64), macOS (amd64/arm64), and Windows (amd64)
- One-line installer script

[Unreleased]: https://github.com/vngcloud/greennode-cli/compare/v1.3.2...HEAD
[1.3.2]: https://github.com/vngcloud/greennode-cli/compare/v1.3.1...v1.3.2
[1.3.1]: https://github.com/vngcloud/greennode-cli/compare/v1.3.0...v1.3.1
[1.3.0]: https://github.com/vngcloud/greennode-cli/compare/v1.2.0...v1.3.0
[1.2.0]: https://github.com/vngcloud/greennode-cli/compare/v1.1.0...v1.2.0
[1.1.0]: https://github.com/vngcloud/greennode-cli/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/vngcloud/greennode-cli/releases/tag/v1.0.0
