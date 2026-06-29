# sshkey list

## Description

List all SSH keys in the current project. Supports pagination and name filtering.

In `table` and `text` output, the public key is shortened to a short preview so
rows stay readable. To see the **full** public key, use `--output json`.

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
ID                                       | NAME        | PUBLIC KEY
-----------------------------------------+-------------+------------------------------------------
ssh-a18340e0-a525-4550-bc7f-30d909cf8447 | minhpq3     | ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC2…
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
