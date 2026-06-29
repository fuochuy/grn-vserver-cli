# floating-ip list

## Description

List all floating IPs (public WAN IP addresses) in the current project. Supports
pagination and name filtering.

`json` output shows the full response. To keep the `table` view compact, the
table format differs slightly:

- columns are shown in a fixed order: `uuid`, `ip`, `networkInterfaceId`,
  `fixedIp`, `status`, `zone`, `createdAt`;
- other fields are hidden;
- the `uuid` and `networkInterfaceId` columns are shortened to a preview;
- the `zone` column shows only the zone **name** instead of the full zone object;
- the `createdAt` timestamp is shown as `DD-MM-YYYY HH:MM`.

Use `--output json` when you need the full values (full `uuid`,
`networkInterfaceId`, the complete zone object, full timestamp, etc.).

## Synopsis

```
grn vserver floating-ip list
    [--name <value>]
    [--page <value>]
    [--page-size <value>]
```

## Options

`--name` (optional)
: Filter results by floating IP name (substring match).

`--page` (default: `1`)
: Page number to retrieve (1-based).

`--page-size` (default: `50`)
: Number of items per page.

## Examples

List all floating IPs:

```bash
grn vserver floating-ip list
```

Get the second page with 10 items per page:

```bash
grn vserver floating-ip list --page 2 --page-size 10
```

Compact table view:

```bash
grn vserver floating-ip list --output table
```

```
UUID                  | IP         | NETWORK INTERFACE ID  | FIXED IP | STATUS   | ZONE   | CREATED AT
----------------------+------------+-----------------------+----------+----------+--------+-----------------
wan-3212ca88-9c7e-46… | 58.84.2.19 | net-in-fa9ada9b-9006… | 10.1.4.4 | ATTACHED | HCM-1A | 29-06-2026 11:21
```

See the full values (including the complete zone object) with JSON output:

```bash
grn vserver floating-ip list --output json
```

Filter output with JMESPath:

```bash
grn vserver floating-ip list --query "listData[*].{ip:ip,fixedIp:fixedIp,status:status}"
```
