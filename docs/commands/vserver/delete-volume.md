# volume delete

## Description

Delete a volume. The command prompts for confirmation before deleting unless `--force` is used.

**This action is irreversible.** Ensure the volume is detached from all servers before deleting.

## Synopsis

```
grn vserver volume delete
    --volume-id <value>
    [--force]
```

## Options

`--volume-id` (required)
: Volume ID to delete.

`--force` (default: `false`)
: Skip the confirmation prompt.

## Examples

Delete a volume (with confirmation):

```bash
grn vserver volume delete --volume-id vol-abc12345-6789-def0-1234-abcdef012345
```

Delete without confirmation:

```bash
grn vserver volume delete --volume-id vol-abc12345-6789-def0-1234-abcdef012345 --force
```
