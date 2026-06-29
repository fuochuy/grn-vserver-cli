# server attach-internal-interface

## Description

Attach an internal network interface to a vServer instance.

The interface is created on the subnet given by `--subnet-id`. Provide `--ip` to
request a specific private IP, or omit it to let the system assign one
automatically (`ip` is sent as `null`).

## Synopsis

```
grn vserver server attach-internal-interface
    --server-id <value>
    --subnet-id <value>
    [--ip <value>]
```

## Options

`--server-id` (required)
: ID of the server. Run [`server list`](list-servers.md) to find it.

`--subnet-id` (required)
: ID of the subnet to create the interface on. Run
[`subnet list`](list-subnets.md) to find it.

`--ip` (optional)
: Private IP to request. If omitted, an IP is assigned automatically.

## Examples

Attach an internal interface with an auto-assigned IP:

```bash
grn vserver server attach-internal-interface \
  --server-id ins-fac86f55-baa6-4f5e-86fd-971aef61cb0e \
  --subnet-id sub-a8b35817-8e5e-40b0-ab22-451ebb75a403
```

Attach an internal interface with a specific IP:

```bash
grn vserver server attach-internal-interface \
  --server-id ins-fac86f55-baa6-4f5e-86fd-971aef61cb0e \
  --subnet-id sub-a8b35817-8e5e-40b0-ab22-451ebb75a403 \
  --ip 10.0.0.5
```

## See also

- [`server detach-internal-interface`](detach-server-internal-interface.md) — detach internal interfaces
- [`subnet list`](list-subnets.md) — list subnets
