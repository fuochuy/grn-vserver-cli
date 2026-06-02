# server list

## Description

List all vServer instances in the current project. Supports pagination and name filtering.

## Synopsis

```
grn vserver server list
    [--name <value>]
    [--page <value>]
    [--page-size <value>]
    [--no-paginate]
```

## Options

`--name` (optional)
: Filter results by server name (substring match, case-insensitive).

`--page` (default: `1`)
: Page number to retrieve (1-based).

`--page-size` (default: `50`)
: Number of items per page.

`--no-paginate` (default: `false`)
: Disable automatic pagination — return only the requested page.

## Examples

List all servers:

```bash
grn vserver server list
```

Filter by name:

```bash
grn vserver server list --name prod
```

Get the second page with 20 items per page:

```bash
grn vserver server list --page 2 --page-size 20
```

Output as table:

```bash
grn vserver server list --output table
```

Filter output with JMESPath:

```bash
grn vserver server list --query "data[*].{id:id,name:name,status:status}"
```
