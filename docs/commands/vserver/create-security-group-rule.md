# secgroup rule create

## Description

Add an inbound (ingress) or outbound (egress) rule to a security group.

## Synopsis

```
grn vserver secgroup rule create
    --secgroup-id <value>
    --direction <value>
    --protocol <value>
    --ether-type <value>
    --remote-ip-prefix <value>
    --port-range-min <value>
    --port-range-max <value>
    [--remote-group-id <value>]
    [--description <value>]
```

## Options

### Required

`--secgroup-id`
: Security group ID to add the rule to.

`--direction`
: Traffic direction. Must be `ingress` (inbound) or `egress` (outbound).

`--protocol`
: Network protocol. Must be one of: `tcp`, `udp`, `icmp`, `any`.

`--ether-type` (default: `IPv4`)
: IP version. Must be `IPv4` or `IPv6`.

`--remote-ip-prefix`
: Remote CIDR block to allow traffic from/to (e.g. `0.0.0.0/0` for all). Required unless `--remote-group-id` is used.

`--port-range-min`
: Minimum port number. Required for `tcp`/`udp`. Must not be set for `icmp` or `any`.

`--port-range-max`
: Maximum port number. Required for `tcp`/`udp`. Must be >= `--port-range-min`.

### Optional

`--remote-group-id`
: Remote security group ID. Traffic is allowed from/to members of this group. Use instead of `--remote-ip-prefix` for group-based rules.

`--description`
: Human-readable rule description.

## Examples

Allow all inbound SSH (port 22) from any IP:

```bash
grn vserver secgroup rule create \
  --secgroup-id secg-52042c19-2706-44db-b38c-2310bb853357 \
  --direction ingress \
  --protocol tcp \
  --ether-type IPv4 \
  --port-range-min 22 \
  --port-range-max 22 \
  --remote-ip-prefix 0.0.0.0/0 \
  --description "Allow SSH"
```

Allow inbound HTTP and HTTPS:

```bash
# HTTP
grn vserver secgroup rule create \
  --secgroup-id secg-52042c19-2706-44db-b38c-2310bb853357 \
  --direction ingress \
  --protocol tcp \
  --ether-type IPv4 \
  --port-range-min 80 \
  --port-range-max 80 \
  --remote-ip-prefix 0.0.0.0/0

# HTTPS
grn vserver secgroup rule create \
  --secgroup-id secg-52042c19-2706-44db-b38c-2310bb853357 \
  --direction ingress \
  --protocol tcp \
  --ether-type IPv4 \
  --port-range-min 443 \
  --port-range-max 443 \
  --remote-ip-prefix 0.0.0.0/0
```

Allow all ICMP (ping) from a specific subnet:

```bash
grn vserver secgroup rule create \
  --secgroup-id secg-52042c19-2706-44db-b38c-2310bb853357 \
  --direction ingress \
  --protocol icmp \
  --ether-type IPv4 \
  --remote-ip-prefix 10.0.0.0/8
```

Allow all outbound traffic:

```bash
grn vserver secgroup rule create \
  --secgroup-id secg-52042c19-2706-44db-b38c-2310bb853357 \
  --direction egress \
  --protocol any \
  --ether-type IPv4 \
  --remote-ip-prefix 0.0.0.0/0
```

Allow traffic from another security group:

```bash
grn vserver secgroup rule create \
  --secgroup-id secg-52042c19-2706-44db-b38c-2310bb853357 \
  --direction ingress \
  --protocol tcp \
  --ether-type IPv4 \
  --port-range-min 3306 \
  --port-range-max 3306 \
  --remote-group-id secg-99887766-5544-3322-1100-aabbccddeeff \
  --description "Allow MySQL from app servers"
```

## See also

- [`secgroup rule list`](list-security-group-rules.md) — list existing rules
- [`secgroup rule delete`](delete-security-group-rule.md) — remove a rule
