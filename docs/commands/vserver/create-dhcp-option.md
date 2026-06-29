# dhcp create

## Description

Create a new DHCP option set.

The two default DNS servers — `10.166.12.196` and `10.166.12.197` — are always
included in a new set. You may add up to **2** more addresses with
`--dns-server`, so a set holds at most **4** DNS servers in total. Passing one
of the defaults again is ignored and does not count toward the limit.

## Synopsis

```
grn vserver dhcp create
    --name <value>
    [--dns-server <ip>]...
```

## Options

`--name` (required)
: Name of the DHCP option set.

`--dns-server` (optional, repeatable)
: An additional DNS server IP address. May be given at most twice (beyond the
two defaults). Each value must be a valid IP address.

## Examples

Create a DHCP option set with only the default DNS servers:

```bash
grn vserver dhcp create --name phanminh211
```

Create a set with two additional DNS servers:

```bash
grn vserver dhcp create \
  --name phanminh211 \
  --dns-server 8.8.8.8 \
  --dns-server 1.1.1.1
```

The resulting `dnsServers` payload is the two defaults followed by the
addresses you add, e.g.
`["10.166.12.196", "10.166.12.197", "8.8.8.8", "1.1.1.1"]`.

## See also

- [`dhcp list`](list-dhcp-options.md) — list DHCP option sets
- [`dhcp get`](get-dhcp-option.md) — show a DHCP option set's details
- [`dhcp delete`](delete-dhcp-option.md) — delete a DHCP option set
