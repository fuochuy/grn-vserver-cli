# network-interface edit

## Description

Rename an elastic network interface. Only the name can be changed.

## Synopsis

```
grn vserver network-interface edit
    --network-interface-id <value>
    --name <value>
```

## Options

`--network-interface-id` (required)
: ID of the network interface to rename. Run
[`network-interface list`](list-network-interfaces.md) to find it.

`--name` (required)
: The new name for the network interface.

## Examples

Rename a network interface:

```bash
grn vserver network-interface edit \
  --network-interface-id network-interface-f2dd56b3-bedc-4ec9-8481-e72c704e82f6 \
  --name minhpq3
```

## See also

- [`network-interface list`](list-network-interfaces.md) — list network interfaces
- [`network-interface delete`](delete-network-interface.md) — delete a network interface
