# vpc create

## Description

Create a new VPC (Virtual Private Cloud / network). Use `--dry-run` to validate the CIDR block without creating the VPC.

## Synopsis

```
grn vserver vpc create
    --name <value>
    --cidr <value>
    [--description <value>]
    [--is-default]
    [--dry-run]
```

## Options

`--name` (required)
: VPC name.

`--cidr` (required)
: CIDR block for the VPC (e.g. `10.0.0.0/16`). Must be a valid IPv4 CIDR.

`--description` (optional)
: Human-readable description of the VPC.

`--is-default` (default: `false`)
: Mark this VPC as the default network for the project.

`--dry-run`
: Validate the name and CIDR without creating the VPC.

## Examples

Create a basic VPC:

```bash
grn vserver vpc create --name my-vpc --cidr 10.0.0.0/16
```

Create the default VPC with a description:

```bash
grn vserver vpc create \
  --name production \
  --cidr 172.16.0.0/12 \
  --description "Production network" \
  --is-default
```

Validate without creating:

```bash
grn vserver vpc create --name my-vpc --cidr 10.0.0.0/16 --dry-run
```

## See also

- [`subnet create`](create-subnet.md) — create subnets inside this VPC
- [`vpc list`](list-vpcs.md) — list existing VPCs
