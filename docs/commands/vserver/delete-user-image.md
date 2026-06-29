# user-image delete

## Description

Delete a single user image by its ID. By default the command prompts for
confirmation before deleting. This action is irreversible.

## Synopsis

```
grn vserver user-image delete
    --user-image-id <value>
    [--force]
```

## Options

`--user-image-id` (required)
: ID of the user image to delete. Run
[`user-image list`](list-user-images.md) to find it.

`--force` (default: `false`)
: Skip the confirmation prompt and delete immediately.

## Examples

Delete a user image (with confirmation prompt):

```bash
grn vserver user-image delete \
  --user-image-id img-fc768036-6346-412f-8cea-c8c18ded7030
```

Delete without the confirmation prompt:

```bash
grn vserver user-image delete \
  --user-image-id img-fc768036-6346-412f-8cea-c8c18ded7030 \
  --force
```

## See also

- [`user-image list`](list-user-images.md) — list available user images
