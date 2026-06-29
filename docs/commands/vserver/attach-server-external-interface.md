# server attach-external-interface

## Description

Attach an existing external (elastic) network interface to a vServer instance.

## Synopsis

```
grn vserver server attach-external-interface
    --server-id <value>
    --network-interface-id <value>
```

## Options

`--server-id` (required)
: ID of the server. Run [`server list`](list-servers.md) to find it.

`--network-interface-id` (required)
: ID of the external (elastic) network interface to attach. Run
[`network-interface list`](list-network-interfaces.md) to find it.

## Examples

Attach an external network interface to a server:

```bash
grn vserver server attach-external-interface \
  --server-id ins-fac86f55-baa6-4f5e-86fd-971aef61cb0e \
  --network-interface-id network-interface-df0021c4-c276-4726-971f-d71101a77dbb
```

## See also

- [`server detach-external-interface`](detach-server-external-interface.md) — detach an external interface
- [`network-interface list`](list-network-interfaces.md) — list network interfaces
