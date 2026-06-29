# network-interface update-tags

## Description

Replace the key/value tag list of an elastic network interface.

The tags you provide become the **complete** list for the interface — any tag
not included is removed. Each tag carries an `isEdited` marker so the backend
can tell which entries were modified:

- `--tag key=value` — a tag whose value was **not** changed (`isEdited=false`);
- `--edited-tag key=value` — a tag whose value **was** changed (`isEdited=true`).

At least one tag (from either flag) is required. The interface's `resourceId`
and the fixed `resourceType` of `NETWORK-INTERFACE` are set for you.

## Synopsis

```
grn vserver network-interface update-tags
    --network-interface-id <value>
    [--tag <key>=<value>]...
    [--edited-tag <key>=<value>]...
```

## Options

`--network-interface-id` (required)
: ID of the network interface whose tags to update. Run
[`network-interface list`](list-network-interfaces.md) to find it.

`--tag` (optional, repeatable)
: An unchanged tag in `key=value` form (sent with `isEdited=false`).

`--edited-tag` (optional, repeatable)
: A changed tag in `key=value` form (sent with `isEdited=true`).

## Examples

Set a single tag that was edited:

```bash
grn vserver network-interface update-tags \
  --network-interface-id network-interface-9d836add-87c7-412a-814d-f9343bda53aa \
  --edited-tag vks-cluster-ids=k8s-b95418e9-30f8-47e9-b756-144e0e0a057c
```

Replace the full list, keeping one tag and editing another:

```bash
grn vserver network-interface update-tags \
  --network-interface-id network-interface-9d836add-87c7-412a-814d-f9343bda53aa \
  --tag env=prod \
  --edited-tag vks-cluster-ids=k8s-b95418e9-30f8-47e9-b756-144e0e0a057c
```

## See also

- [`network-interface list`](list-network-interfaces.md) — find the interface ID
- [`network-interface create`](create-network-interface.md) — create a network interface
