# subnet get

## Description

Get details of a specific subnet by its ID and the parent VPC ID.

## Synopsis

```
grn vserver subnet get
    --subnet-id <value>
    --vpc-id <value>
```

## Options

`--subnet-id` (required)
: Subnet ID (e.g. `sub-0305c2cc-a89f-4cd3-bdd0-49cd312049a6`).

`--vpc-id` (required)
: VPC (network) ID the subnet belongs to.

## Examples

Get subnet details:

```bash
grn vserver subnet get \
  --subnet-id sub-0305c2cc-a89f-4cd3-bdd0-49cd312049a6 \
  --vpc-id net-83f3e9a9-71e6-438d-89e8-bde44722d987
```
