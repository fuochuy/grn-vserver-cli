# sshkey delete

## Description

Delete a single SSH key by its ID. By default the command prompts for
confirmation before deleting. This action is irreversible.

## Synopsis

```
grn vserver sshkey delete
    --sshkey-id <value>
    [--force]
```

## Options

`--sshkey-id` (required)
: ID of the SSH key to delete. Run [`sshkey list`](list-sshkeys.md) to find it.

`--force` (default: `false`)
: Skip the confirmation prompt and delete immediately.

## Examples

Delete an SSH key (with confirmation prompt):

```bash
grn vserver sshkey delete --sshkey-id ssh-a18340e0-a525-4550-bc7f-30d909cf8447
```

Delete without the confirmation prompt:

```bash
grn vserver sshkey delete \
  --sshkey-id ssh-a18340e0-a525-4550-bc7f-30d909cf8447 \
  --force
```

## See also

- [`sshkey list`](list-sshkeys.md) — list available SSH keys
- [`sshkey create`](create-sshkey.md) — create a new SSH key pair
- [`sshkey import`](import-sshkey.md) — import an existing public key
