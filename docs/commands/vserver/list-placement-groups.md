# placement-group list

## Description

List all placement groups (server groups) in the current project. Supports
pagination and name filtering.

`json` output shows the full response. To keep the `table` view compact, the
table format differs slightly:

- columns are shown in a fixed order: `uuid`, `name`, `policyName`,
  `description`, `servers`, `createdAt`;
- other fields (e.g. `policyId`, `serverGroupId`) are hidden;
- the `uuid` column is shortened to a preview;
- the `servers` column shows only the server **names** (comma-separated) instead
  of the full nested objects.

Use `--output json` when you need the full values (full `uuid`, `policyId`, etc.).

## Synopsis

```
grn vserver placement-group list
    [--name <value>]
    [--page <value>]
    [--page-size <value>]
```

## Options

`--name` (optional)
: Filter results by placement group name (substring match).

`--page` (default: `1`)
: Page number to retrieve (1-based).

`--page-size` (default: `50`)
: Number of items per page.

## Examples

List all placement groups:

```bash
grn vserver placement-group list
```

Filter by name:

```bash
grn vserver placement-group list --name web
```

Get the second page with 10 items per page:

```bash
grn vserver placement-group list --page 2 --page-size 10
```

Output as a table (compact columns, server names only):

```bash
grn vserver placement-group list --output table
```

```
UUID                   | NAME | POLICY NAME        | DESCRIPTION | SERVERS                             | CREATED AT
-----------------------+------+--------------------+-------------+-------------------------------------+-----------
server-group-4f9904a…  | pg-a | SOFT ANTI AFFINITY | my group    | mp-pfsense_pfSense-CE-2_8_1, web-01 | 2026-06-29
```

See the full values (including `policyId` and full `uuid`) with JSON output:

```bash
grn vserver placement-group list --output json
```

Filter output with JMESPath:

```bash
grn vserver placement-group list --query "listData[*].{uuid:uuid,name:name,policyId:policyId}"
```

## See also

- [`placement-group delete`](delete-placement-group.md) — delete a placement group
