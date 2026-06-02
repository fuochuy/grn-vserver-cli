# secgroup rule list

## Description

List all rules in a security group.

## Synopsis

```
grn vserver secgroup rule list
    --secgroup-id <value>
    [--page <value>]
    [--page-size <value>]
```

## Options

`--secgroup-id` (required)
: Security group ID whose rules to list.

`--page` (default: `1`)
: Page number (1-based).

`--page-size` (default: `50`)
: Number of items per page.

## Examples

List rules in a security group:

```bash
grn vserver secgroup rule list --secgroup-id secg-52042c19-2706-44db-b38c-2310bb853357
```

Output as table:

```bash
grn vserver secgroup rule list \
  --secgroup-id secg-52042c19-2706-44db-b38c-2310bb853357 \
  --output table
```
