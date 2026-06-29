# dhcp delete

## Description

Delete a single DHCP option set by its ID. By default the command prompts for
confirmation before deleting. This action is irreversible.

## Synopsis

```
grn vserver dhcp delete
    --dhcp-option-id <value>
    [--force]
```

## Options

`--dhcp-option-id` (required)
: ID of the DHCP option set to delete. Run [`dhcp list`](list-dhcp-options.md)
to find it.

`--force` (default: `false`)
: Skip the confirmation prompt and delete immediately.

## Examples

Delete a DHCP option set (with confirmation prompt):

```bash
grn vserver dhcp delete --dhcp-option-id dop-17c25106-04e6-4839-8961-408f8828b796
```

Delete without the confirmation prompt:

```bash
grn vserver dhcp delete \
  --dhcp-option-id dop-17c25106-04e6-4839-8961-408f8828b796 \
  --force
```

## See also

- [`dhcp list`](list-dhcp-options.md) — list available DHCP option sets
