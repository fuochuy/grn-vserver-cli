# sshkey import

## Description

Import an existing SSH **public** key so it can be attached to servers. Unlike
[`sshkey create`](create-sshkey.md), no private key is generated — you supply a
public key you already own.

To keep the CLI experience clean, prefer `--public-key-file` to read the key
straight from a `.pub` file (e.g. `~/.ssh/id_rsa.pub`) instead of pasting the
long key string on the command line. Provide **exactly one** of
`--public-key-file` or `--public-key`.

## Synopsis

```
grn vserver sshkey import
    --name <value>
    (--public-key-file <path> | --public-key <value>)
```

## Options

`--name` (required)
: Name for the imported SSH key.

`--public-key-file` (optional)
: Path to a public key file to read (e.g. `~/.ssh/id_rsa.pub`). The file
contents are sent as the public key. Recommended over `--public-key`.

`--public-key` (optional)
: The public key contents inline (e.g. `ssh-rsa AAAA...`). Use only when you
cannot point to a file.

The value is whitespace-trimmed and must start with a recognised SSH public key
type (`ssh-rsa`, `ssh-ed25519`, `ecdsa-sha2-...`, etc.).

The printed response shows only a preview of the public key. Use
`sshkey list --output json` to see the full key.

## Examples

Import from a public key file (recommended):

```bash
grn vserver sshkey import --name minhtesting --public-key-file ~/.ssh/id_rsa.pub
```

Import by pasting the key inline:

```bash
grn vserver sshkey import \
  --name minhtesting \
  --public-key "ssh-rsa AAAAB3NzaC1yc2E... lap15005@host"
```

## See also

- [`sshkey create`](create-sshkey.md) — generate a brand-new key pair
- [`sshkey list`](list-sshkeys.md) — list existing SSH keys
- [`sshkey delete`](delete-sshkey.md) — delete an SSH key
