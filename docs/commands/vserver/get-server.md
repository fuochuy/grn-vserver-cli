# server get

## Description

Get details of a specific vServer instance by its ID.

## Synopsis

```
grn vserver server get
    --server-id <value>
```

## Options

`--server-id` (required)
: ID of the vServer instance to retrieve (e.g. `srv-abc12345-6789-def0-1234-abcdef012345`).

## Examples

Get server details:

```bash
grn vserver server get --server-id srv-abc12345-6789-def0-1234-abcdef012345
```

Output as JSON:

```bash
grn vserver server get --server-id srv-abc12345-6789-def0-1234-abcdef012345 --output json
```

Extract specific fields with JMESPath:

```bash
grn vserver server get \
  --server-id srv-abc12345-6789-def0-1234-abcdef012345 \
  --query "data.{id:id,name:name,status:status,ip:addresses}"
```
