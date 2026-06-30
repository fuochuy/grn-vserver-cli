# server list-interfaces

## Description

List the network interfaces attached to a vServer instance.

Interfaces are returned in two groups: **internal** (PRIVATE) and **external**
(PUBLIC). In table output each group is rendered as its own table, prefixed with
a count; the other formats (`json`, `text`) return the full response unchanged.

## Synopsis

```
grn vserver server list-interfaces
    --server-id <value>
```

## Options

`--server-id` (required)
: ID of the server whose interfaces to list. Run [`server list`](list-servers.md)
to find it.

## Examples

List the interfaces of a server:

```bash
grn vserver server list-interfaces \
  --server-id ins-fac86f55-baa6-4f5e-86fd-971aef61cb0e
```

Render the result as two tables:

```bash
grn vserver server list-interfaces \
  --server-id ins-fac86f55-baa6-4f5e-86fd-971aef61cb0e \
  --output table
```

```
Internal interfaces (2):

Uuid                 | Fixed Ip   | Floating Ip | Status | Interface Type | Subnet Uuid          | Mac               | Created At
---------------------+------------+-------------+--------+----------------+----------------------+-------------------+-----------------
net-in-95f91833-8d38… | 10.24.10.3 |             | ACTIVE | PRIVATE        | sub-0fdc68a1-662e-48… | 02:19:83:81:b8:bf | 29-06-2026 23:47
net-in-aeb3fb0d-6349… | 10.24.11.3 |             | ACTIVE | PRIVATE        | sub-a8b35817-8e5e-40… | 02:81:84:08:12:0d | 30-06-2026 09:16

External interfaces (0):

(none)
```

## See also

- [`server attach-internal-interface`](attach-server-internal-interface.md) — attach an internal interface
- [`server attach-external-interface`](attach-server-external-interface.md) — attach an external interface
- [`server detach-internal-interface`](detach-server-internal-interface.md) — detach internal interfaces
- [`server detach-external-interface`](detach-server-external-interface.md) — detach an external interface
