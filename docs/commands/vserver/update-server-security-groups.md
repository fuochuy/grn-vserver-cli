# server update-secgroup

## Description

Update the security groups attached to a vServer instance.

The security group IDs you provide become the **complete** set for the server —
any security group not listed is detached. Run [`secgroup list`](list-security-groups.md)
to see available security groups.

## Synopsis

```
grn vserver server update-secgroup
    --server-id <value>
    --security-group <value>
```

## Options

`--server-id` (required)
: ID of the server whose security groups will be updated.

`--security-group` (required)
: Security group IDs to attach, comma-separated. This replaces the server's
current security groups. Each ID must be valid (run `secgroup list` to discover them).

## Examples

Attach a single security group to a server:

```bash
grn vserver server update-secgroup \
  --server-id ins-0ee2e3f2-3955-48c5-9483-ce80d0766387 \
  --security-group secg-52042c19-2706-44db-b38c-2310bb853357
```

Replace a server's security groups with multiple groups:

```bash
grn vserver server update-secgroup \
  --server-id ins-0ee2e3f2-3955-48c5-9483-ce80d0766387 \
  --security-group secg-52042c19-2706-44db-b38c-2310bb853357,secg-47cbfccd-7c36-4b94-93da-0c638407a182
```

Output as table:

```bash
grn vserver server update-secgroup \
  --server-id ins-0ee2e3f2-3955-48c5-9483-ce80d0766387 \
  --security-group secg-52042c19-2706-44db-b38c-2310bb853357 \
  --output table
```

## See also

- [`secgroup list`](list-security-groups.md) — list available security groups
- [`server get`](get-server.md) — view a server's current configuration
