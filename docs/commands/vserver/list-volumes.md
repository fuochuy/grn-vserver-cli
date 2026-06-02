# volume list

## Description

List all volumes in the current project. Supports pagination and name filtering.

## Synopsis

```
grn vserver volume list
    [--name <value>]
    [--page <value>]
    [--page-size <value>]
```

## Options

`--name` (optional)
: Filter by volume name (substring match).

`--page` (default: `1`)
: Page number (1-based).

`--page-size` (default: `50`)
: Number of items per page.

## Examples

List all volumes:

```bash
grn vserver volume list
```

Filter by name:

```bash
grn vserver volume list --name data
```

Output as table:

```bash
grn vserver volume list --output table
```

Extract volume IDs and sizes:

```bash
grn vserver volume list --query "data[*].{id:id,name:name,size:size,status:status}"
```
