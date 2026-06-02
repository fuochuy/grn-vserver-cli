# secgroup get

## Description

Get details of a specific security group by its ID.

## Synopsis

```
grn vserver secgroup get
    --secgroup-id <value>
```

## Options

`--secgroup-id` (required)
: Security group ID (e.g. `secg-52042c19-2706-44db-b38c-2310bb853357`).

## Examples

Get security group details:

```bash
grn vserver secgroup get --secgroup-id secg-52042c19-2706-44db-b38c-2310bb853357
```

Output as JSON:

```bash
grn vserver secgroup get --secgroup-id secg-52042c19-2706-44db-b38c-2310bb853357 --output json
```
