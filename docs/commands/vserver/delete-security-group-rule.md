# secgroup rule delete

## Description

Remove a rule from a security group.

## Synopsis

```
grn vserver secgroup rule delete
    --secgroup-id <value>
    --rule-id <value>
```

## Options

`--secgroup-id` (required)
: Security group ID the rule belongs to.

`--rule-id` (required)
: Security group rule ID to delete.

## Examples

Delete a rule:

```bash
grn vserver secgroup rule delete \
  --secgroup-id secg-52042c19-2706-44db-b38c-2310bb853357 \
  --rule-id rule-abc12345-6789-def0-1234-abcdef012345
```

## See also

- [`secgroup rule list`](list-security-group-rules.md) — find the rule ID to delete
