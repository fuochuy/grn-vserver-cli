# volume get

## Description

Get details of a specific volume by its ID.

## Synopsis

```
grn vserver volume get
    --volume-id <value>
```

## Options

`--volume-id` (required)
: Volume ID (e.g. `vol-abc12345-6789-def0-1234-abcdef012345`).

## Examples

Get volume details:

```bash
grn vserver volume get --volume-id vol-abc12345-6789-def0-1234-abcdef012345
```

Output as JSON:

```bash
grn vserver volume get --volume-id vol-abc12345-6789-def0-1234-abcdef012345 --output json
```

Extract specific fields:

```bash
grn vserver volume get \
  --volume-id vol-abc12345-6789-def0-1234-abcdef012345 \
  --query "data.{id:id,name:name,size:size,volumeTypeId:volumeTypeId,status:status}"
```
