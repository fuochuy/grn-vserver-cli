# dhcp list

## Description

List all DHCP option sets in the current project. Supports pagination and name
filtering.

`json` output shows the full response. To keep the `table` view readable, the
table format differs slightly:

- columns are shown in a fixed order: `uuid`, `name`, `status`, `dnsServers`,
  `associatedVpcs`, `createdAt`;
- other fields are hidden;
- the `uuid` column is shortened to a preview;
- the `createdAt` timestamp is shown as `DD-MM-YYYY HH:MM`;
- `associatedVpcs` is a derived count — the number of entries in the option's
  `associatedNetworks` list;
- each **DNS server** is shown on its own line: the first address appears on the
  option's main row, and each additional address adds an extra row in which only
  the `dnsServers` column is filled.

Use `--output json` to see the full values (full `uuid`, the complete
`dnsServers` array, `associatedNetworks`, full timestamp, etc.).

## Synopsis

```
grn vserver dhcp list
    [--name <value>]
    [--page <value>]
    [--page-size <value>]
```

## Options

`--name` (optional)
: Filter results by DHCP option name (substring match).

`--page` (default: `1`)
: Page number to retrieve (1-based).

`--page-size` (default: `50`)
: Number of items per page.

## Examples

List all DHCP options:

```bash
grn vserver dhcp list
```

Compact table view (DNS servers expanded across rows):

```bash
grn vserver dhcp list --output table
```

```
UUID                  | NAME   | STATUS | DNS SERVERS   | ASSOCIATED VPCS | CREATED AT
----------------------+--------+--------+---------------+-----------------+-----------------
dop-17c25106-04e6-48… | dhcp-a | ACTIVE | 10.19.255.211 | 2               | 29-06-2026 11:21
                      |        |        | 10.19.255.227 |                 |
                      |        |        | 10.19.255.243 |                 |
```

See the full values with JSON output:

```bash
grn vserver dhcp list --output json
```

## See also

- [`dhcp delete`](delete-dhcp-option.md) — delete a DHCP option set
