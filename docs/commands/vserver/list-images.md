# image list

## Description

List available OS or GPU images that can be used when creating a vServer instance.

## Synopsis

```
grn vserver image list
    --type <value>
    [--image-version <value>]
    [--page <value>]
    [--page-size <value>]
```

## Options

`--type` (required)
: Image type. Must be one of:
  - `os` — standard operating system images (Ubuntu, CentOS, Windows, etc.)
  - `gpu` — GPU-optimised images with CUDA/driver pre-installed

`--image-version` (optional)
: Filter by image version string (substring match, case-insensitive). Useful for finding a specific OS release, e.g. `22.04` for Ubuntu 22.04.

`--page` (default: `1`)
: Page number (1-based).

`--page-size` (default: `50`)
: Number of items per page.

## Examples

List all OS images:

```bash
grn vserver image list --type os
```

List GPU images:

```bash
grn vserver image list --type gpu
```

Filter for Ubuntu 22.04:

```bash
grn vserver image list --type os --image-version 22.04
```

Output as table:

```bash
grn vserver image list --type os --output table
```
