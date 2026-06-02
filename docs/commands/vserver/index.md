# vServer Commands

vServer (VNG Virtual Server) commands for managing virtual machine instances, volumes, networking, and security groups.

```bash
grn vserver <resource> <command> [options]
```

## Available commands

### Server

| Command | Description |
|---------|-------------|
| [server list](list-servers.md) | List all vServer instances |
| [server get](get-server.md) | Get vServer instance details |
| [server create](create-server.md) | Create a new vServer instance |
| [server start](start-server.md) | Start a stopped vServer instance |
| [server stop](stop-server.md) | Stop a running vServer instance |
| [server reboot](reboot-server.md) | Reboot a vServer instance |
| [server resize](resize-server.md) | Resize a vServer instance to a different flavor |
| [server delete](delete-server.md) | Delete a vServer instance |

### Volume

| Command | Description |
|---------|-------------|
| [volume list](list-volumes.md) | List all volumes |
| [volume get](get-volume.md) | Get volume details |
| [volume create](create-volume.md) | Create a new volume |
| [volume resize](resize-volume.md) | Resize a volume's size or change its volume type |
| [volume delete](delete-volume.md) | Delete a volume |

### VPC (Network)

| Command | Description |
|---------|-------------|
| [vpc list](list-vpcs.md) | List all VPCs |
| [vpc get](get-vpc.md) | Get VPC details |
| [vpc create](create-vpc.md) | Create a new VPC |
| [vpc delete](delete-vpc.md) | Delete a VPC |

### Subnet

| Command | Description |
|---------|-------------|
| [subnet list](list-subnets.md) | List subnets in a VPC |
| [subnet get](get-subnet.md) | Get subnet details |
| [subnet create](create-subnet.md) | Create a new subnet |
| [subnet delete](delete-subnet.md) | Delete a subnet |

### Security Group

| Command | Description |
|---------|-------------|
| [secgroup list](list-security-groups.md) | List all security groups |
| [secgroup get](get-security-group.md) | Get security group details |
| [secgroup create](create-security-group.md) | Create a new security group |
| [secgroup delete](delete-security-group.md) | Delete a security group |
| [secgroup rule list](list-security-group-rules.md) | List rules in a security group |
| [secgroup rule get](get-security-group-rule.md) | Get a security group rule |
| [secgroup rule create](create-security-group-rule.md) | Add a rule to a security group |
| [secgroup rule delete](delete-security-group-rule.md) | Remove a rule from a security group |

### Reference

| Command | Description |
|---------|-------------|
| [flavor list](list-flavors.md) | List available flavors (CPU/RAM) |
| [image list](list-images.md) | List available OS and GPU images |
| [volume-type list](list-volume-types.md) | List available volume types for a zone |
