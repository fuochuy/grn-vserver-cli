# server detach-internal-interface

## Description

Detach one or more internal network interfaces from a vServer instance.

## Synopsis

```
grn vserver server detach-internal-interface
    --server-id <value>
    --network-interface-id <value>
```

## Options

`--server-id` (required)
: ID of the server. Run [`server list`](list-servers.md) to find it.

`--network-interface-id` (required)
: Internal network interface IDs to detach, comma-separated. At least one is
required.

## Examples

Detach a single internal interface:

```bash
grn vserver server detach-internal-interface \
  --server-id ins-fac86f55-baa6-4f5e-86fd-971aef61cb0e \
  --network-interface-id net-in-1e712a8b-1878-44bd-b197-02540a5ea89a
```

Detach multiple internal interfaces at once:

```bash
grn vserver server detach-internal-interface \
  --server-id ins-fac86f55-baa6-4f5e-86fd-971aef61cb0e \
  --network-interface-id net-in-1e712a8b-1878-44bd-b197-02540a5ea89a,net-in-80109a77-2679-44fb-9910-a783612d546a
```

## See also

- [`server attach-internal-interface`](attach-server-internal-interface.md) — attach an internal interface
