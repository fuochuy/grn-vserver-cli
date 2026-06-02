# server resize

## Description

Resize a vServer instance by changing it to a different flavor (CPU/RAM profile). The server is stopped before resizing and automatically restarted after.

Use `grn vserver flavor list` to find available flavor IDs.

## Synopsis

```
grn vserver server resize
    --server-id <value>
    --flavor-id <value>
```

## Options

`--server-id` (required)
: ID of the vServer instance to resize.

`--flavor-id` (required)
: New flavor ID. Run `grn vserver flavor list` to see available flavors for your zone.

## Examples

Resize to a larger flavor:

```bash
grn vserver server resize \
  --server-id srv-abc12345-6789-def0-1234-abcdef012345 \
  --flavor-id flav-dca36eca-2018-4607-bde4-ac9d9cd31655
```

## See also

- [`server get`](get-server.md) — verify instance status after resize
- [`flavor list`](list-flavors.md) — discover available flavors
