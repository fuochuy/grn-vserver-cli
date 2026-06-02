# secgroup create

## Description

Create a new security group. After creation, add inbound and outbound rules with `grn vserver secgroup rule create`.

## Synopsis

```
grn vserver secgroup create
    --name <value>
    [--description <value>]
```

## Options

`--name` (required)
: Security group name.

`--description` (optional)
: Human-readable description.

## Examples

Create a security group:

```bash
grn vserver secgroup create --name web-servers
```

Create with a description:

```bash
grn vserver secgroup create \
  --name web-servers \
  --description "Allow HTTP/HTTPS inbound from anywhere"
```

## See also

- [`secgroup rule create`](create-security-group-rule.md) — add rules to this security group
- [`server create`](create-server.md) — attach a security group with `--security-group`
