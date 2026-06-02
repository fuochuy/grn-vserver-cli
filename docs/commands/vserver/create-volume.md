# volume create

## Description

Create a new block storage volume. Use `--dry-run` to validate parameters without creating the volume.

## Synopsis

```
grn vserver volume create
    --name <value>
    --volume-type-id <value>
    --size <value>
    --zone-id <value>
    [--description <value>]
    [--encryption-type <value>]
    [--multiattach]
    [--is-poc]
    [--dry-run]
```

## Options

### Required

`--name`
: Volume name.

`--volume-type-id`
: Volume type ID. Run `grn vserver volume-type list` to find available types.

`--size`
: Volume size in GiB.

`--zone-id`
: Availability zone ID (e.g. `HCM03-1A`). Omit to see available zones.

### Optional

`--description`
: Human-readable description.

`--encryption-type`
: Encryption type to apply to the volume.

`--multiattach` (default: `false`)
: Allow the volume to be attached to multiple servers simultaneously.

`--is-poc` (default: `false`)
: Mark as a proof-of-concept (PoC) volume for billing purposes.

`--dry-run`
: Validate all parameters without creating the volume.

## Examples

Create a 50 GiB SSD volume:

```bash
grn vserver volume create \
  --name my-data-volume \
  --volume-type-id vtype-7a7a8610-34f5-11ee-be56-0242ac120002 \
  --size 50 \
  --zone-id HCM03-1A
```

Create a multi-attach volume:

```bash
grn vserver volume create \
  --name shared-volume \
  --volume-type-id vtype-7a7a8610-34f5-11ee-be56-0242ac120002 \
  --size 100 \
  --zone-id HCM03-1A \
  --multiattach \
  --description "Shared storage for cluster"
```

Validate without creating:

```bash
grn vserver volume create \
  --name my-data-volume \
  --volume-type-id vtype-7a7a8610-34f5-11ee-be56-0242ac120002 \
  --size 50 \
  --zone-id HCM03-1A \
  --dry-run
```

## See also

- [`volume-type list`](list-volume-types.md) — find available volume types
- [`volume resize`](resize-volume.md) — expand a volume after creation
