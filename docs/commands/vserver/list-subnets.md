# subnet list

## Description

List all subnets within a VPC.

## Synopsis

```
grn vserver subnet list
    --vpc-id <value>
    [--page <value>]
    [--page-size <value>]
```

## Options

`--vpc-id` (required)
: VPC (network) ID whose subnets to list.

`--page` (default: `1`)
: Page number (1-based).

`--page-size` (default: `50`)
: Number of items per page.

## Examples

List subnets in a VPC:

```bash
grn vserver subnet list --vpc-id net-83f3e9a9-71e6-438d-89e8-bde44722d987
```

Output as table:

```bash
grn vserver subnet list --vpc-id net-83f3e9a9-71e6-438d-89e8-bde44722d987 --output table
```

Extract subnet IDs:

```bash
grn vserver subnet list \
  --vpc-id net-83f3e9a9-71e6-438d-89e8-bde44722d987 \
  --query "data[*].{id:id,name:name,cidr:cidr}"
```
