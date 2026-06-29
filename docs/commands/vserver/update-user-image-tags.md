# user-image update-tags

## Description

Replace the key/value tag list of a user image.

The tags you provide become the **complete** list for the image — any tag not
included is removed. Each tag carries an `isEdited` marker so the backend can
tell which entries were modified:

- `--tag key=value` — a tag whose value was **not** changed (`isEdited=false`);
- `--edited-tag key=value` — a tag whose value **was** changed (`isEdited=true`).

At least one tag (from either flag) is required. The image's `resourceId` and
the fixed `resourceType` of `IMAGE` are set for you.

## Synopsis

```
grn vserver user-image update-tags
    --user-image-id <value>
    [--tag <key>=<value>]...
    [--edited-tag <key>=<value>]...
```

## Options

`--user-image-id` (required)
: ID of the user image whose tags to update. Run
[`user-image list`](list-user-images.md) to find it.

`--tag` (optional, repeatable)
: An unchanged tag in `key=value` form (sent with `isEdited=false`).

`--edited-tag` (optional, repeatable)
: A changed tag in `key=value` form (sent with `isEdited=true`).

## Examples

Set a single tag that was edited:

```bash
grn vserver user-image update-tags \
  --user-image-id img-de53bd9b-b93b-4104-b850-088b24a78b2e \
  --edited-tag vks-cluster-ids=k8s-b95418e9-30f8-47e9-b756-144e0e0a057c
```

Replace the full list, keeping one tag and editing another:

```bash
grn vserver user-image update-tags \
  --user-image-id img-de53bd9b-b93b-4104-b850-088b24a78b2e \
  --tag env=prod \
  --edited-tag vks-cluster-ids=k8s-b95418e9-30f8-47e9-b756-144e0e0a057c
```

## See also

- [`user-image list`](list-user-images.md) — find the image ID
- [`server tag-key`](list-tag-keys.md) — list available tag keys
- [`server tag-value`](list-tag-values.md) — list values for a tag key
