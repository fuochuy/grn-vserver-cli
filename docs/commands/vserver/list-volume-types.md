# volume-type list

## Description

List available volume types for a given availability zone. Use this to find the `--volume-type-id` / `--root-disk-type-id` required when creating servers or volumes.

The command requires two pieces of information:

1. `--zone-id` — the availability zone you want to provision in.
2. `--type` — the volume type zone name (e.g. `SSD`, `NVMe`). If omitted, the command prints the available type names for the zone.

## Synopsis

```
grn vserver volume-type list
    --zone-id <value>
    [--type <value>]
```

## Options

`--zone-id` (required)
: Availability zone ID (e.g. `HCM03-1A`). Omit to see a list of available zones.

`--type` (optional)
: Volume type zone name to filter by (e.g. `SSD`, `NVMe`). If not provided, the command lists the available type names for the zone.

## Examples

Discover available type names for a zone (omit `--type`):

```bash
grn vserver volume-type list --zone-id HCM03-1A
```

List all SSD volume types:

```bash
grn vserver volume-type list --zone-id HCM03-1A --type SSD
```

Output as table:

```bash
grn vserver volume-type list --zone-id HCM03-1A --type SSD --output table
```

## See also

- [`volume create`](create-volume.md) — uses `--volume-type-id` from this list
- [`server create`](create-server.md) — uses `--root-disk-type-id` from this list
