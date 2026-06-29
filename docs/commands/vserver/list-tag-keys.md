# server tag-key

## Description

List every tag key that exists in the current project.

Tag keys are paired with values to label resources — for example when creating
an image with [`server create-image`](create-server-image.md). Use
[`server tag-value`](list-tag-values.md) to see the values recorded for a
specific key.

## Synopsis

```
grn vserver server tag-key
```

## Options

This command takes no options beyond the global flags (`--output`, `--query`,
`--profile`, etc.).

## Examples

List all tag keys:

```bash
grn vserver server tag-key
```

List tag keys as a table:

```bash
grn vserver server tag-key --output table
```

## See also

- [`server tag-value`](list-tag-values.md) — list the values for a tag key
- [`server create-image`](create-server-image.md) — create an image with tags
