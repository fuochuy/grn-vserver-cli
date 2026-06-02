# secgroup delete

## Description

Delete a security group. The command prompts for confirmation before deleting unless `--force` is used.

**This action is irreversible.** Detach the security group from all servers before deleting it.

## Synopsis

```
grn vserver secgroup delete
    --secgroup-id <value>
    [--force]
```

## Options

`--secgroup-id` (required)
: Security group ID to delete.

`--force` (default: `false`)
: Skip the confirmation prompt.

## Examples

Delete a security group (with confirmation):

```bash
grn vserver secgroup delete --secgroup-id secg-52042c19-2706-44db-b38c-2310bb853357
```

Delete without confirmation:

```bash
grn vserver secgroup delete --secgroup-id secg-52042c19-2706-44db-b38c-2310bb853357 --force
```
