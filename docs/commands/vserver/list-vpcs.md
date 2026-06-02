# vpc list

## Description

List all VPCs (Virtual Private Clouds / networks) in the current project.

## Synopsis

```
grn vserver vpc list
    [--name <value>]
    [--page <value>]
    [--page-size <value>]
```

## Options

`--name` (optional)
: Filter by VPC name (substring match).

`--page` (default: `1`)
: Page number (1-based).

`--page-size` (default: `50`)
: Number of items per page.

## Examples

List all VPCs:

```bash
grn vserver vpc list
```

Filter by name:

```bash
grn vserver vpc list --name production
```

Output as table:

```bash
grn vserver vpc list --output table
```

Extract VPC IDs with JMESPath:

```bash
grn vserver vpc list --query "data[*].{id:id,name:name,cidr:cidr}"
```
