# network-interface list

## Description

List all elastic network interfaces in the current project. Supports pagination
and name filtering.

`json` output shows the full response. To keep the `table` view compact, the
table format differs slightly:

- columns are shown in a fixed order: `uuid`, `name`, `status`, `vpcName`,
  `serverName`, `ip`, `zone`, `createdAt`;
- other fields are hidden;
- the `uuid` column is shortened to a preview;
- the `zone` column shows only the zone **name** instead of the full zone object;
- the `createdAt` timestamp is shown as `DD-MM-YYYY HH:MM`.

Use `--output json` when you need the full values (full `uuid`, the complete zone
object, full timestamp, etc.).

## Synopsis

```
grn vserver network-interface list
    [--name <value>]
    [--page <value>]
    [--page-size <value>]
```

## Options

`--name` (optional)
: Filter results by network interface name (substring match).

`--page` (default: `1`)
: Page number to retrieve (1-based).

`--page-size` (default: `50`)
: Number of items per page.

## Examples

List all network interfaces:

```bash
grn vserver network-interface list
```

Filter by name:

```bash
grn vserver network-interface list --name nic
```

Compact table view:

```bash
grn vserver network-interface list --output table
```

```
UUID                  | NAME  | STATUS | VPC NAME | SERVER NAME | IP       | ZONE   | CREATED AT
----------------------+-------+--------+----------+-------------+----------+--------+-----------------
network-interface-f2… | nic-1 | ACTIVE | my-vpc   | web-01      | 10.1.4.4 | HCM-1A | 29-06-2026 11:21
```

See the full values with JSON output:

```bash
grn vserver network-interface list --output json
```

## See also

- [`network-interface edit`](edit-network-interface.md) — rename a network interface
- [`network-interface delete`](delete-network-interface.md) — delete a network interface
