# dhcp list-vpcs

## Description

List every VPC (network) currently associated with a DHCP option set.

`json` output shows the full response. For `table` output the `id` column is
shortened to a preview and timestamps are shown as `DD-MM-YYYY HH:MM`; the
columns are `id`, `displayName`, `cidr`, `status`, `createdAt`.

## Synopsis

```
grn vserver dhcp list-vpcs
    --dhcp-option-id <value>
    [--page <value>]
    [--page-size <value>]
```

## Options

`--dhcp-option-id` (required)
: ID of the DHCP option set whose associated VPCs to list. Run
[`dhcp list`](list-dhcp-options.md) to find it.

`--page` (default: `1`)
: Page number to retrieve (1-based).

`--page-size` (default: `50`)
: Number of items per page.

## Examples

List the VPCs associated with a DHCP option set:

```bash
grn vserver dhcp list-vpcs --dhcp-option-id dop-2c2c0d85-32e8-41fa-a8f2-c685e81b2bef
```

Compact table view:

```bash
grn vserver dhcp list-vpcs \
  --dhcp-option-id dop-2c2c0d85-32e8-41fa-a8f2-c685e81b2bef \
  --output table
```

## See also

- [`dhcp get`](get-dhcp-option.md) — show a DHCP option set's details
- [`dhcp associate-vpc`](associate-dhcp-vpc.md) — associate or detach a VPC
