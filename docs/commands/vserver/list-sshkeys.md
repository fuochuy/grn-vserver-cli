# sshkey list

## Description

List all SSH keys in the current project. Supports pagination and name filtering.

In `table` output the columns are shown in a fixed, compact order — `id`,
`name`, `status`, `createdAt`, `pubKey` — with the long public key last. The
`id` is shortened to a preview, the public key is truncated, and the timestamp is
shown as `DD-MM-YYYY HH:MM`. To see the **full** values (full id, full public
key, full timestamp), use `--output json`.

## Synopsis

```
grn vserver sshkey list
    [--name <value>]
    [--page <value>]
    [--page-size <value>]
```

## Options

`--name` (optional)
: Filter results by SSH key name (substring match).

`--page` (default: `1`)
: Page number to retrieve (1-based).

`--page-size` (default: `50`)
: Number of items per page.

## Examples

List all SSH keys:

```bash
grn vserver sshkey list
```

Filter by name:

```bash
grn vserver sshkey list --name deploy
```

Get the second page with 5 items per page:

```bash
grn vserver sshkey list --page 2 --page-size 5
```

Output as a table (public key shown as a short preview):

```bash
grn vserver sshkey list --output table
```

```
ID                    | NAME     | STATUS | CREATED AT       | PUB KEY
----------------------+----------+--------+------------------+------------------------------------------
ssh-c5532395-c242-4b… | sshkey-1 | ACTIVE | 29-06-2026 16:02 | ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDE…
```

See the full public key with JSON output:

```bash
grn vserver sshkey list --output json
```

Filter output with JMESPath:

```bash
grn vserver sshkey list --query "listData[*].{id:id,name:name}"
```

## See also

- [`sshkey create`](create-sshkey.md) — create a new SSH key pair
- [`sshkey import`](import-sshkey.md) — import an existing public key
- [`sshkey delete`](delete-sshkey.md) — delete an SSH key
