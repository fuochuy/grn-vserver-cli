# vpc delete

## Description

Delete a VPC. The command prompts for confirmation before deleting unless `--force` is used.

**This action is irreversible.** Ensure all subnets and resources within the VPC are removed before deleting it.

## Synopsis

```
grn vserver vpc delete
    --vpc-id <value>
    [--force]
```

## Options

`--vpc-id` (required)
: VPC (network) ID to delete.

`--force` (default: `false`)
: Skip the confirmation prompt.

## Examples

Delete a VPC (with confirmation):

```bash
grn vserver vpc delete --vpc-id net-83f3e9a9-71e6-438d-89e8-bde44722d987
```

Delete without confirmation:

```bash
grn vserver vpc delete --vpc-id net-83f3e9a9-71e6-438d-89e8-bde44722d987 --force
```
