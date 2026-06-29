# placement-group delete

## Description

Delete a single placement group (server group) by its ID. By default the command
prompts for confirmation before deleting. This action is irreversible.

## Synopsis

```
grn vserver placement-group delete
    --placement-group-id <value>
    [--force]
```

## Options

`--placement-group-id` (required)
: ID of the placement group to delete. Run
[`placement-group list`](list-placement-groups.md) to find it.

`--force` (default: `false`)
: Skip the confirmation prompt and delete immediately.

## Examples

Delete a placement group (with confirmation prompt):

```bash
grn vserver placement-group delete \
  --placement-group-id server-group-4f9904a1-158c-4bd4-99f7-9eb4d51635a3
```

Delete without the confirmation prompt:

```bash
grn vserver placement-group delete \
  --placement-group-id server-group-4f9904a1-158c-4bd4-99f7-9eb4d51635a3 \
  --force
```

## See also

- [`placement-group list`](list-placement-groups.md) — list available placement groups
