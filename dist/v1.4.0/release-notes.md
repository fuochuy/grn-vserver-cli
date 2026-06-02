### Added
- Full command reference documentation for all vServer resources: server, volume, VPC, subnet, security group, security group rules, images, flavors, and volume types (26 new pages under `docs/commands/vserver/`)
- `CHANGELOG.md`, `CONTRIBUTING.md`, and `CONTRIBUTORS.md`
- Release tooling: `scripts/release`, `scripts/bump-version`, `scripts/render-changelog`, `scripts/extract-release-notes`
- GitHub Actions CI workflow (`ci.yml`) — runs `go vet`, `go test`, and build on every PR and push to main
- GitHub Actions release workflow (`release.yml`) — builds all platform binaries, generates `checksums.txt`, extracts release notes from `CHANGELOG.md`, and publishes a GitHub release on tag push
- Makefile targets: `bump-patch`, `bump-minor`, `bump-major`, `release`, `release-dry-run`
- `.changes/` changelog fragment system for per-PR changelog entries

---
