# network-interface create

## Description

Create a new elastic network interface in a zone.

Tags are optional key/value pairs that label the interface. Created interfaces
appear under [`network-interface list`](list-network-interfaces.md).

## Synopsis

```
grn vserver network-interface create
    --name <value>
    --zone-id <value>
    [--tag <key>=<value>]...
```

## Options

`--name` (required)
: Name of the network interface.

`--zone-id` (required)
: Availability zone ID to create the interface in (for example `HCM03-1C`).

`--tag` (optional, repeatable)
: A tag to attach, in `key=value` form. Pass the flag once per pair, for example
`--tag env=prod --tag vks-cluster-ids=k8s-...`.

## Examples

Create a network interface:

```bash
grn vserver network-interface create \
  --name network-interface \
  --zone-id HCM03-1C
```

Create one with a tag:

```bash
grn vserver network-interface create \
  --name network-interface \
  --zone-id HCM03-1C \
  --tag vks-cluster-ids=k8s-b95418e9-30f8-47e9-b756-144e0e0a057c
```

## See also

- [`network-interface list`](list-network-interfaces.md) — list network interfaces
- [`network-interface update-tags`](update-network-interface-tags.md) — update an interface's tags
