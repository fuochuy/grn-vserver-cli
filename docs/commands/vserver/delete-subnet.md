# subnet delete

## Description

Delete a subnet from a VPC. The command prompts for confirmation before deleting unless `--force` is used.

**This action is irreversible.**

## Synopsis

```
grn vserver subnet delete
    --subnet-id <value>
    --vpc-id <value>
    [--force]
```

## Options

`--subnet-id` (required)
: Subnet ID to delete.

`--vpc-id` (required)
: VPC (network) ID the subnet belongs to.

`--force` (default: `false`)
: Skip the confirmation prompt.

## Examples

Delete a subnet (with confirmation):

```bash
grn vserver subnet delete \
  --subnet-id sub-0305c2cc-a89f-4cd3-bdd0-49cd312049a6 \
  --vpc-id net-83f3e9a9-71e6-438d-89e8-bde44722d987
```

Delete without confirmation:

```bash
grn vserver subnet delete \
  --subnet-id sub-0305c2cc-a89f-4cd3-bdd0-49cd312049a6 \
  --vpc-id net-83f3e9a9-71e6-438d-89e8-bde44722d987 \
  --force
```
