# server create-image

## Description

Create a new user image (custom image) from an existing vServer instance.

The image is created from the server identified by `--server-id` and named with
`--name`. You can optionally attach tags — key/value pairs used to label the
image. Created images appear under [`user-image list`](list-user-images.md).

## Synopsis

```
grn vserver server create-image
    --server-id <value>
    --name <value>
    [--tag <key>=<value>]...
```

## Options

`--server-id` (required)
: ID of the server to create the image from. Run [`server list`](list-servers.md)
to find it.

`--name` (required)
: Name of the new image.

`--tag` (optional, repeatable)
: A tag to attach, in `key=value` form. Pass the flag once per pair, for example
`--tag env=prod --tag team=infra`. Run [`server tag-key`](list-tag-keys.md) and
[`server tag-value`](list-tag-values.md) to discover existing keys and values.

## Examples

Create an image from a server:

```bash
grn vserver server create-image \
  --server-id ins-c0481b83-c7a7-460e-8008-72676f0ca51c \
  --name minhpq3
```

Create an image with tags:

```bash
grn vserver server create-image \
  --server-id ins-c0481b83-c7a7-460e-8008-72676f0ca51c \
  --name minhpq3 \
  --tag env=prod \
  --tag team=infra
```

## See also

- [`server list`](list-servers.md) — find the source server ID
- [`server tag-key`](list-tag-keys.md) — list available tag keys
- [`server tag-value`](list-tag-values.md) — list values for a tag key
- [`user-image list`](list-user-images.md) — list created user images
