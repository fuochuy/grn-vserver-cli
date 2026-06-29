# server tag-value

## Description

List every value that has been used for a given tag key in the project.

Run [`server tag-key`](list-tag-keys.md) first to discover the available keys.

## Synopsis

```
grn vserver server tag-value
    --key <value>
```

## Options

`--key` (required)
: Tag key whose values to list. Run [`server tag-key`](list-tag-keys.md) to find
valid keys.

## Examples

List the values recorded for the `vks-cluster-ids` key:

```bash
grn vserver server tag-value --key vks-cluster-ids
```

List the values as a table:

```bash
grn vserver server tag-value --key vks-cluster-ids --output table
```

## See also

- [`server tag-key`](list-tag-keys.md) — list available tag keys
- [`server create-image`](create-server-image.md) — create an image with tags
