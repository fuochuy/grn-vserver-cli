# placement-group create

## Description

Create a new placement group. A placement **policy** is required.

Because policy IDs are hard to remember, you usually omit `--policy-id`: the
command then lists the available policies and prompts you to pick one by number.
If you already know the policy ID, pass `--policy-id` to create the group
non-interactively.

## Synopsis

```
grn vserver placement-group create
    --name <value>
    [--description <value>]
    [--policy-id <value>]
```

## Options

`--name` (required)
: Name for the new placement group.

`--description` (optional)
: Description of the placement group.

`--policy-id` (optional)
: Placement policy ID. Run
[`placement-group list-policies`](list-placement-group-policies.md) to see the
options. If omitted, you are prompted to choose a policy interactively.

## Examples

Create with interactive policy selection (no `--policy-id`):

```bash
grn vserver placement-group create --name web-tier --description "frontend nodes"
```

```
Available policies:
  [1] SOFT ANTI AFFINITY    aaa7d316-cff2-11eb-b8bc-0242ac130003
  [2] SOFT AFFINITY         a2162216-cff2-11eb-b8bc-0242ac130003
Select a policy by number [1-2]: 1
```

Create non-interactively with a known policy ID:

```bash
grn vserver placement-group create \
  --name web-tier \
  --description "frontend nodes" \
  --policy-id aaa7d316-cff2-11eb-b8bc-0242ac130003
```

## See also

- [`placement-group list-policies`](list-placement-group-policies.md) — discover policy IDs
- [`placement-group edit`](edit-placement-group.md) — update a placement group
- [`placement-group list`](list-placement-groups.md) — list existing placement groups
