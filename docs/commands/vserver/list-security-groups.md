# secgroup list

## Description

List all security groups in the current project.

## Synopsis

```
grn vserver secgroup list
    [--name <value>]
    [--page <value>]
    [--page-size <value>]
```

## Options

`--name` (optional)
: Filter by security group name (substring match).

`--page` (default: `1`)
: Page number (1-based).

`--page-size` (default: `50`)
: Number of items per page.

## Examples

List all security groups:

```bash
grn vserver secgroup list
```

Filter by name:

```bash
grn vserver secgroup list --name web
```

Output as table:

```bash
grn vserver secgroup list --output table
```

Extract IDs:

```bash
grn vserver secgroup list --query "data[*].{id:id,name:name}"
```
