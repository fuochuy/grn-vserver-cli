# vpc get

## Description

Get details of a specific VPC (Virtual Private Cloud / network) by its ID.

## Synopsis

```
grn vserver vpc get
    --vpc-id <value>
```

## Options

`--vpc-id` (required)
: VPC (network) ID (e.g. `net-83f3e9a9-71e6-438d-89e8-bde44722d987`).

## Examples

Get VPC details:

```bash
grn vserver vpc get --vpc-id net-83f3e9a9-71e6-438d-89e8-bde44722d987
```

Output as JSON:

```bash
grn vserver vpc get --vpc-id net-83f3e9a9-71e6-438d-89e8-bde44722d987 --output json
```
