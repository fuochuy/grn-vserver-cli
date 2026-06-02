# subnet create

## Description

Create a new subnet within an existing VPC.

## Synopsis

```
grn vserver subnet create
    --vpc-id <value>
    --cidr <value>
    --zone-id <value>
    [--name <value>]
```

## Options

`--vpc-id` (required)
: VPC (network) ID to create the subnet in.

`--cidr` (required)
: CIDR block for the subnet (e.g. `10.0.1.0/24`). Must be a subset of the parent VPC's CIDR.

`--zone-id` (required)
: Availability zone ID (e.g. `HCM03-1A`). Omit to see available zones.

`--name` (optional)
: Subnet name.

## Examples

Create a subnet:

```bash
grn vserver subnet create \
  --vpc-id net-83f3e9a9-71e6-438d-89e8-bde44722d987 \
  --cidr 10.0.1.0/24 \
  --zone-id HCM03-1A \
  --name web-tier
```

Create without a name:

```bash
grn vserver subnet create \
  --vpc-id net-83f3e9a9-71e6-438d-89e8-bde44722d987 \
  --cidr 10.0.2.0/24 \
  --zone-id HCM03-1A
```

## See also

- [`vpc list`](list-vpcs.md) — find the VPC ID
- [`subnet list`](list-subnets.md) — verify the subnet was created
