# server detach-external-interface

## Description

Detach an external (elastic) network interface from a vServer instance.

## Synopsis

```
grn vserver server detach-external-interface
    --server-id <value>
    --network-interface-id <value>
```

## Options

`--server-id` (required)
: ID of the server. Run [`server list`](list-servers.md) to find it.

`--network-interface-id` (required)
: ID of the external (elastic) network interface to detach. Run
[`network-interface list`](list-network-interfaces.md) to find it.

## Examples

Detach an external network interface from a server:

```bash
grn vserver server detach-external-interface \
  --server-id ins-fac86f55-baa6-4f5e-86fd-971aef61cb0e \
  --network-interface-id net-in-80109a77-2679-44fb-9910-a783612d546a
```

## See also

- [`server attach-external-interface`](attach-server-external-interface.md) — attach an external interface
- [`network-interface list`](list-network-interfaces.md) — list network interfaces
