# sshkey create

## Description

Create a new SSH key pair. The server generates **both** the public and the
private key. The private key is returned **only once**, so the CLI saves it
automatically as a `<name>.pem` file in your **Downloads** directory.

An existing file is never overwritten — if `<name>.pem` already exists, a
`(n)` suffix is added (e.g. `minhpq3 (1).pem`), the same way browsers handle
duplicate downloads. The file is written with `0600` permissions so only your
user can read it. Use `--output-dir` to save it somewhere other than Downloads.

The printed response shows only a **preview** of the public and private keys to
keep the output readable. The full private key is in the saved `.pem` file; the
full public key can be retrieved later with `sshkey list --output json`.

## Synopsis

```
grn vserver sshkey create
    --name <value>
    [--output-dir <path>]
```

## Options

`--name` (required)
: Name for the new SSH key. Also used as the filename of the saved
`<name>.pem` private key.

`--output-dir` (optional)
: Directory in which to save the `<name>.pem` private key. Defaults to your
Downloads directory. The directory is created if it does not exist.

## Examples

Create a new key pair:

```bash
grn vserver sshkey create --name minhpq3
```

The private key is saved to `~/Downloads/minhpq3.pem` and the path is printed:

```
Private key saved to: /home/you/Downloads/minhpq3.pem
Keep this file safe — the private key cannot be retrieved again.
Run 'grn vserver sshkey list --output json' to see the full public key.
```

Save the private key to a specific directory:

```bash
grn vserver sshkey create --name minhpq3 --output-dir ~/.ssh
```

Use the saved key to SSH into a server:

```bash
chmod 600 ~/Downloads/minhpq3.pem
ssh -i ~/Downloads/minhpq3.pem <user>@<server-ip>
```

## See also

- [`sshkey import`](import-sshkey.md) — import an existing public key instead
- [`sshkey list`](list-sshkeys.md) — list existing SSH keys
- [`sshkey delete`](delete-sshkey.md) — delete an SSH key
