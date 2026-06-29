# dhcp get

## Description

Show the full details of a single DHCP option set by its ID.

`json` output shows the complete response. For `table` output the view mirrors
[`dhcp list`](list-dhcp-options.md): the `uuid` is shortened, the timestamp is
shown as `DD-MM-YYYY HH:MM`, `associatedVpcs` is the count of entries in
`associatedNetworks`, and each DNS server is shown on its own row.

## Synopsis

```
grn vserver dhcp get
    --dhcp-option-id <value>
```

## Options

`--dhcp-option-id` (required)
: ID of the DHCP option set. Run [`dhcp list`](list-dhcp-options.md) to find it.

## Examples

Get a DHCP option set:

```bash
grn vserver dhcp get --dhcp-option-id dop-2c2c0d85-32e8-41fa-a8f2-c685e81b2bef
```

Compact table view:

```bash
grn vserver dhcp get \
  --dhcp-option-id dop-2c2c0d85-32e8-41fa-a8f2-c685e81b2bef \
  --output table
```

## See also

- [`dhcp list`](list-dhcp-options.md) — list DHCP option sets
- [`dhcp list-vpcs`](list-dhcp-vpcs.md) — list VPCs associated with a DHCP option set
