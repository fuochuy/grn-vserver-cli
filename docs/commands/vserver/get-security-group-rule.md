# secgroup rule get

## Description

Get details of a specific rule within a security group.

## Synopsis

```
grn vserver secgroup rule get
    --secgroup-id <value>
    --rule-id <value>
```

## Options

`--secgroup-id` (required)
: Security group ID the rule belongs to.

`--rule-id` (required)
: Security group rule ID.

## Examples

Get rule details:

```bash
grn vserver secgroup rule get \
  --secgroup-id secg-52042c19-2706-44db-b38c-2310bb853357 \
  --rule-id rule-abc12345-6789-def0-1234-abcdef012345
```
