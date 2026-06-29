# user-image list

## Description

List all user images in the current project. User images are custom images you
have created from your servers — distinct from the public OS/GPU images shown by
[`image list`](list-images.md). Supports pagination and name filtering.

`json` output shows the full response. To keep the `table` view compact, the
table format differs slightly:

- columns are shown in a fixed order: `uuid`, `name`, `minDisk`, `imageSize`,
  `status`, `createdAt`;
- other fields are hidden;
- the `uuid` column is shortened to a preview;
- the `createdAt` timestamp is shown as `DD-MM-YYYY HH:MM`.

Use `--output json` when you need the full values (full `uuid`, full timestamp,
and any other fields).

## Synopsis

```
grn vserver user-image list
    [--name <value>]
    [--page <value>]
    [--page-size <value>]
```

## Options

`--name` (optional)
: Filter results by user image name (substring match).

`--page` (default: `1`)
: Page number to retrieve (1-based).

`--page-size` (default: `50`)
: Number of items per page.

## Examples

List all user images:

```bash
grn vserver user-image list
```

Filter by name:

```bash
grn vserver user-image list --name backup
```

Get the second page with 10 items per page:

```bash
grn vserver user-image list --page 2 --page-size 10
```

Compact table view:

```bash
grn vserver user-image list --output table
```

```
UUID                  | NAME        | MIN DISK | IMAGE SIZE      | STATUS | CREATED AT
----------------------+-------------+----------+-----------------+--------+-----------------
img-fc768036-6346-41… | my-snapshot | 20       | 1.073741824e+10 | ACTIVE | 29-06-2026 16:02
```

See the full values with JSON output:

```bash
grn vserver user-image list --output json
```

## See also

- [`user-image delete`](delete-user-image.md) — delete a user image
- [`image list`](list-images.md) — list public OS and GPU images
