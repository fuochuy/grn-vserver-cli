# placement-group list-policies

## Description

List the placement group policies that can be used when creating a placement
group (see [`placement-group create`](create-placement-group.md)).

The API returns each policy description in both English (`description`) and
Vietnamese (`descriptionVi`). The CLI shows a single description in the language
you choose with `--language` (English by default).

In `table` output the description is omitted to keep rows compact: only `uuid`,
`name`, and `status` are shown (in that order), and the `uuid` is shortened to a
preview. Use `--output json` to see the full values.

## Synopsis

```
grn vserver placement-group list-policies
    [--language <en|vi>]
```

## Options

`--language` (default: `en`)
: Language of the policy description: `en` (English) or `vi` (Vietnamese).

## Examples

List policies (English descriptions):

```bash
grn vserver placement-group list-policies
```

List policies with Vietnamese descriptions:

```bash
grn vserver placement-group list-policies --language vi
```

Compact table view (`uuid`, `name`, `status`):

```bash
grn vserver placement-group list-policies --output table
```

```
UUID                   | NAME              | STATUS
-----------------------+-------------------+--------
aaa7d316-cff2-11eb-b… | SOFT ANTI AFFINITY | ACTIVE
a2162216-cff2-11eb-b… | SOFT AFFINITY      | ACTIVE
```

## See also

- [`placement-group create`](create-placement-group.md) — create a placement group using a policy
- [`placement-group list`](list-placement-groups.md) — list existing placement groups
