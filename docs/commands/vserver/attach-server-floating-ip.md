# server attach-floating-ip

## Description

Attach a floating IP (WAN IP) to a network interface of a vServer instance.

The floating IP is attached to the network interface identified by
`--network-interface-id` on the given server.

## Synopsis

```
grn vserver server attach-floating-ip
    --server-id <value>
    --floating-ip-id <value>
    --network-interface-id <value>
```

## Options

`--server-id` (required)
: ID of the server. Run [`server list`](list-servers.md) to find it.

`--floating-ip-id` (required)
: ID of the floating IP to attach. Run [`floating-ip list`](list-floating-ips.md)
to find it.

`--network-interface-id` (required)
: ID of the network interface to attach the floating IP to. Run
[`network-interface list`](list-network-interfaces.md) to find it.

## Examples

Attach a floating IP to a server's network interface:

```bash
grn vserver server attach-floating-ip \
  --server-id ins-fac86f55-baa6-4f5e-86fd-971aef61cb0e \
  --floating-ip-id wan-acd1c96c-42df-4929-802a-1374ccae90f2 \
  --network-interface-id net-in-1e712a8b-1878-44bd-b197-02540a5ea89a
```

## See also

- [`server detach-floating-ip`](detach-server-floating-ip.md) — detach a floating IP
- [`floating-ip list`](list-floating-ips.md) — list floating IPs
- [`network-interface list`](list-network-interfaces.md) — list network interfaces
