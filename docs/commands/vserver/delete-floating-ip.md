# floating-ip delete

## Description

Delete a single floating IP (WAN IP) by its ID. By default the command prompts
for confirmation before deleting. This action is irreversible.

## Synopsis

```
grn vserver floating-ip delete
    --floating-ip-id <value>
    [--force]
```

## Options

`--floating-ip-id` (required)
: ID of the floating IP to delete. Run [`floating-ip list`](list-floating-ips.md)
to find it.

`--force` (default: `false`)
: Skip the confirmation prompt and delete immediately.

## Examples

Delete a floating IP (with confirmation prompt):

```bash
grn vserver floating-ip delete --floating-ip-id wan-f9807aff-4fee-47e3-8ceb-dc267bf9ddb3
```

Delete without the confirmation prompt:

```bash
grn vserver floating-ip delete \
  --floating-ip-id wan-f9807aff-4fee-47e3-8ceb-dc267bf9ddb3 \
  --force
```

## See also

- [`floating-ip list`](list-floating-ips.md) — list available floating IPs
