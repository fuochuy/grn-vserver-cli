# network-interface delete

## Description

Delete a single elastic network interface by its ID. By default the command
prompts for confirmation before deleting. This action is irreversible.

## Synopsis

```
grn vserver network-interface delete
    --network-interface-id <value>
    [--force]
```

## Options

`--network-interface-id` (required)
: ID of the network interface to delete. Run
[`network-interface list`](list-network-interfaces.md) to find it.

`--force` (default: `false`)
: Skip the confirmation prompt and delete immediately.

## Examples

Delete a network interface (with confirmation prompt):

```bash
grn vserver network-interface delete \
  --network-interface-id network-interface-f2dd56b3-bedc-4ec9-8481-e72c704e82f6
```

Delete without the confirmation prompt:

```bash
grn vserver network-interface delete \
  --network-interface-id network-interface-f2dd56b3-bedc-4ec9-8481-e72c704e82f6 \
  --force
```

## See also

- [`network-interface list`](list-network-interfaces.md) — list network interfaces
- [`network-interface edit`](edit-network-interface.md) — rename a network interface
