# Changelog fragments

Add one file per pull request. The file name pattern is `<type>-<short-description>.md`.

Valid types: `feature`, `fix`, `change`, `docs`, `internal`

Example (`feature-subnet-commands.md`):

```markdown
---
type: feature
---

Added `grn vserver subnet` commands: `list`, `get`, `create`, `delete`.
```

These fragments are merged into the top-level `CHANGELOG.md` when a release is cut with `scripts/bump-version`.
