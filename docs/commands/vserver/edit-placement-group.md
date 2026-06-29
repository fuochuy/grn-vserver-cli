# placement-group edit

## Description

Update a placement group's name and/or description. Only the flags you provide
are changed — omit a flag to leave that field untouched. The placement group's
policy cannot be changed with this command.

## Synopsis

```
grn vserver placement-group edit
    --placement-group-id <value>
    [--name <value>]
    [--description <value>]
```

## Options

`--placement-group-id` (required)
: ID of the placement group to update. Run
[`placement-group list`](list-placement-groups.md) to find it.

`--name` (optional)
: New name. Omit to keep the current name.

`--description` (optional)
: New description. Omit to keep the current description.

At least one of `--name` or `--description` must be provided.

## Examples

Rename a placement group:

```bash
grn vserver placement-group edit \
  --placement-group-id server-group-82d97bcf-e582-466c-b363-137462f90049 \
  --name minhphanquang
```

Update both name and description:

```bash
grn vserver placement-group edit \
  --placement-group-id server-group-82d97bcf-e582-466c-b363-137462f90049 \
  --name minhphanquang \
  --description "testing new update"
```

## See also

- [`placement-group create`](create-placement-group.md) — create a placement group
- [`placement-group list`](list-placement-groups.md) — list existing placement groups
