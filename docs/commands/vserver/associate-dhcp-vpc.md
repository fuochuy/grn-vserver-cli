# dhcp associate-vpc

## Description

Associate a VPC (network) with a DHCP option set, or detach it.

A VPC can belong to **only one** DHCP option set. This command updates the VPC's
DHCP association:

- with `--dhcp-option-id`, the VPC is associated with that set (replacing any
  current association);
- with `--detach`, the VPC's association is removed entirely (the VPC then uses
  no custom DHCP option set).

Exactly one of `--dhcp-option-id` or `--detach` must be given.

## Synopsis

```
grn vserver dhcp associate-vpc
    --vpc-id <value>
    (--dhcp-option-id <value> | --detach)
```

## Options

`--vpc-id` (required)
: ID of the VPC (network) to update. Run [`vpc list`](list-vpcs.md) to find it.

`--dhcp-option-id`
: DHCP option set ID to associate the VPC with. Run
[`dhcp list`](list-dhcp-options.md) to find it.

`--detach` (default: `false`)
: Remove the VPC's association with its current DHCP option set.

## Examples

Associate a VPC with a DHCP option set:

```bash
grn vserver dhcp associate-vpc \
  --vpc-id net-d8edff08-6345-4e4c-8acc-16618e8e32db \
  --dhcp-option-id dop-2c2c0d85-32e8-41fa-a8f2-c685e81b2bef
```

Detach a VPC from its DHCP option set:

```bash
grn vserver dhcp associate-vpc \
  --vpc-id net-d8edff08-6345-4e4c-8acc-16618e8e32db \
  --detach
```

## See also

- [`dhcp list-vpcs`](list-dhcp-vpcs.md) — list VPCs associated with a DHCP option set
- [`dhcp get`](get-dhcp-option.md) — show a DHCP option set's details
