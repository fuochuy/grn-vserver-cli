package mcp

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"
)

// McpCmd starts the MCP server over stdio.
var McpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "Start the MCP server for AI assistant integration",
	Long: `Start a Model Context Protocol (MCP) server that exposes grn vserver
commands as tools for AI assistants (Claude Code, Claude Desktop).

Setup for Claude Code — add to ~/.claude.json or .claude/mcp.json:

  {
    "mcpServers": {
      "greennode": {
        "command": "grn",
        "args": ["mcp"]
      }
    }
  }

Then restart Claude Code. You can then ask Claude to manage your vServer
resources in natural language.`,
	RunE: runMCP,
}

func runMCP(cmd *cobra.Command, args []string) error {
	s := server.NewMCPServer("greennode-vserver", "1.0.0",
		server.WithToolCapabilities(true),
	)

	registerDiscoveryTools(s)
	registerServerTools(s)
	registerVolumeTools(s)
	registerNetworkTools(s)
	registerSecgroupTools(s)
	registerSSHKeyTools(s)
	registerPlacementGroupTools(s)
	registerUserImageTools(s)
	registerFloatingIPTools(s)
	registerNetworkInterfaceTools(s)
	registerDhcpTools(s)

	return server.ServeStdio(s)
}

// ── helpers ───────────────────────────────────────────────────────────────────

// grn runs the grn binary with the given arguments and returns stdout+stderr.
func grn(args ...string) (string, error) {
	bin, err := os.Executable()
	if err != nil {
		bin = "grn"
	}
	// Always request JSON output for machine-readable results
	fullArgs := append(args, "--output", "json")
	out, err := exec.Command(bin, fullArgs...).CombinedOutput()
	return string(out), err
}

func textResult(output string, err error) *mcp.CallToolResult {
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error: %s\nOutput: %s", err.Error(), output))
	}
	return mcp.NewToolResultText(strings.TrimSpace(output))
}

func opt(s string) mcp.ToolOption {
	return mcp.WithString(s, mcp.Description(s))
}
func req(s, desc string) mcp.ToolOption {
	return mcp.WithString(s, mcp.Required(), mcp.Description(desc))
}

func getArgs(r mcp.CallToolRequest) map[string]any {
	return r.GetArguments()
}

func sarg(a map[string]any, key string) string {
	v, _ := a[key].(string)
	return v
}

// ── discovery tools ───────────────────────────────────────────────────────────

func registerDiscoveryTools(s *server.MCPServer) {
	// zones
	s.AddTool(mcp.NewTool("list_zones",
		mcp.WithDescription("List all available availability zones. Call this first to find valid --zone-id values."),
	), func(_ context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		out, err := grn("vserver", "volume-type", "list")
		// We parse the stderr suggestion output; just return combined output
		return textResult(out, err), nil
	})

	// images
	s.AddTool(mcp.NewTool("list_images",
		mcp.WithDescription("List available server images. Use imageType 'os' for Linux/Windows images, 'gpu' for GPU-optimized images. Returns image IDs needed for server creation."),
		req("imageType", "Image type: 'os' or 'gpu'"),
		opt("imageVersion"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		a := getArgs(r)
		cmdArgs := []string{"vserver", "image", "list", "--type", sarg(a, "imageType")}
		if v := sarg(a, "imageVersion"); v != "" {
			cmdArgs = append(cmdArgs, "--image-version", v)
		}
		return textResult(grn(cmdArgs...)), nil
	})

	// flavor families
	s.AddTool(mcp.NewTool("list_flavor_families",
		mcp.WithDescription("List available instance families (e.g. general-purpose, compute-optimized, gpu). Use this before list_flavors to find valid family names."),
	), func(_ context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return textResult(grn("vserver", "flavor", "list-families")), nil
	})

	// flavor codes
	s.AddTool(mcp.NewTool("list_flavor_codes",
		mcp.WithDescription("List available CPU platform codes (e.g. code-g, code-a40). Use this before list_flavors to find valid code values."),
	), func(_ context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return textResult(grn("vserver", "flavor", "list-codes")), nil
	})

	// flavors
	s.AddTool(mcp.NewTool("list_flavors",
		mcp.WithDescription("List available flavors (CPU/RAM configurations) for a given instance family and CPU platform code. Returns flavor IDs needed for server creation or resize. Only shows flavors with available capacity."),
		req("family", "Instance family name (from list_flavor_families)"),
		req("code", "CPU platform code (from list_flavor_codes)"),
		opt("zoneId"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		a := getArgs(r)
		cmdArgs := []string{"vserver", "flavor", "list",
			"--family", sarg(a, "family"),
			"--code", sarg(a, "code"),
		}
		if z := sarg(a, "zoneId"); z != "" {
			cmdArgs = append(cmdArgs, "--zone-id", z)
		}
		return textResult(grn(cmdArgs...)), nil
	})

	// volume types
	s.AddTool(mcp.NewTool("list_volume_types",
		mcp.WithDescription("List available volume types for a zone (e.g. SSD, NVMe). Returns volume type IDs needed for server and volume creation."),
		req("zoneId", "Availability zone ID (from list_zones)"),
		req("volumeType", "Volume type zone name: SSD or NVMe"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		a := getArgs(r)
		return textResult(grn("vserver", "volume-type", "list",
			"--zone-id", sarg(a, "zoneId"),
			"--type", sarg(a, "volumeType"),
		)), nil
	})
}

// ── server tools ───────────────────────────────────────────────────────────────

func registerServerTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool("list_servers",
		mcp.WithDescription("List all vServer instances in the project."),
		opt("name"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		cmdArgs := []string{"vserver", "server", "list"}
		if n := sarg(getArgs(r), "name"); n != "" {
			cmdArgs = append(cmdArgs, "--name", n)
		}
		return textResult(grn(cmdArgs...)), nil
	})

	s.AddTool(mcp.NewTool("get_server",
		mcp.WithDescription("Get details of a specific vServer instance including status, IP addresses, and configuration."),
		req("serverId", "Server ID"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return textResult(grn("vserver", "server", "get",
			"--server-id", sarg(getArgs(r), "serverId"),
		)), nil
	})

	s.AddTool(mcp.NewTool("create_server",
		mcp.WithDescription("Create a new vServer instance. Before calling this, use list_zones, list_images, list_flavor_families + list_flavors, list_volume_types, list_vpcs, and list_subnets to find valid IDs. Set attachFloating=true to assign a public floating IP to the server so it is reachable from the internet."),
		req("name", "Server name (alphanumeric, hyphens, underscores; 5-65 chars)"),
		req("zoneId", "Availability zone ID"),
		req("networkId", "VPC ID"),
		req("subnetId", "Subnet ID"),
		req("imageId", "Image ID"),
		req("flavorId", "Flavor ID"),
		req("rootDiskTypeId", "Root disk volume type ID"),
		mcp.WithNumber("rootDiskSize", mcp.Description("Root disk size in GiB (minimum 20, default 20)")),
		mcp.WithBoolean("attachFloating", mcp.Description("Attach a floating (public) IP to the server. When true, GreenNode automatically allocates and associates a public IP address, making the server directly reachable from the internet. Default false.")),
		opt("securityGroup"),
		opt("sshKeyId"),
		opt("userName"),
		opt("userPassword"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		a := getArgs(r)
		cmdArgs := []string{"vserver", "server", "create",
			"--name", sarg(a, "name"),
			"--zone-id", sarg(a, "zoneId"),
			"--network-id", sarg(a, "networkId"),
			"--subnet-id", sarg(a, "subnetId"),
			"--image-id", sarg(a, "imageId"),
			"--flavor-id", sarg(a, "flavorId"),
			"--root-disk-type-id", sarg(a, "rootDiskTypeId"),
		}
		if sz, ok := a["rootDiskSize"].(float64); ok && sz > 0 {
			cmdArgs = append(cmdArgs, "--root-disk-size", fmt.Sprintf("%d", int(sz)))
		}
		if v, _ := a["attachFloating"].(bool); v {
			cmdArgs = append(cmdArgs, "--attach-floating")
		}
		if v := sarg(a, "securityGroup"); v != "" {
			cmdArgs = append(cmdArgs, "--security-group", v)
		}
		if v := sarg(a, "sshKeyId"); v != "" {
			cmdArgs = append(cmdArgs, "--ssh-key-id", v)
		}
		if v := sarg(a, "userName"); v != "" {
			cmdArgs = append(cmdArgs, "--user-name", v)
		}
		if v := sarg(a, "userPassword"); v != "" {
			cmdArgs = append(cmdArgs, "--user-password", v)
		}
		return textResult(grn(cmdArgs...)), nil
	})

	for _, action := range []struct{ name, desc, endpoint string }{
		{"start_server", "Start a stopped vServer instance.", "start"},
		{"stop_server", "Stop a running vServer instance.", "stop"},
		{"reboot_server", "Reboot a vServer instance.", "reboot"},
	} {
		action := action
		s.AddTool(mcp.NewTool(action.name,
			mcp.WithDescription(action.desc),
			req("serverId", "Server ID"),
		), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return textResult(grn("vserver", "server", action.endpoint,
				"--server-id", getArgs(r)["serverId"].(string),
			)), nil
		})
	}

	s.AddTool(mcp.NewTool("resize_server",
		mcp.WithDescription("Resize a vServer instance to a different flavor (CPU/RAM). Use list_flavors to find a valid flavorId first."),
		req("serverId", "Server ID"),
		req("flavorId", "New flavor ID (from list_flavors)"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return textResult(grn("vserver", "server", "resize",
			"--server-id", getArgs(r)["serverId"].(string),
			"--flavor-id", getArgs(r)["flavorId"].(string),
		)), nil
	})

	s.AddTool(mcp.NewTool("update_server_secgroups",
		mcp.WithDescription("Replace the set of security groups attached to a vServer instance. The provided list becomes the complete set — any security group not included is detached. Use list_secgroups to find valid security group IDs."),
		req("serverId", "Server ID"),
		req("securityGroups", "Security group IDs to attach (comma-separated)"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		a := getArgs(r)
		return textResult(grn("vserver", "server", "update-secgroup",
			"--server-id", sarg(a, "serverId"),
			"--security-group", sarg(a, "securityGroups"),
		)), nil
	})

	s.AddTool(mcp.NewTool("delete_server",
		mcp.WithDescription("Delete a vServer instance. This action is irreversible. Always confirm with the user before calling this."),
		req("serverId", "Server ID"),
		mcp.WithBoolean("deleteAllVolumes", mcp.Description("Also delete all volumes attached to the server (default false)")),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := []string{"vserver", "server", "delete", "--force",
			"--server-id", getArgs(r)["serverId"].(string),
		}
		if v, _ := getArgs(r)["deleteAllVolumes"].(bool); v {
			args = append(args, "--delete-all-volumes")
		}
		return textResult(grn(args...)), nil
	})

	s.AddTool(mcp.NewTool("list_tag_keys",
		mcp.WithDescription("List all tag keys available in the project. Tag keys are used with create_server_image."),
	), func(_ context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return textResult(grn("vserver", "server", "tag-key")), nil
	})

	s.AddTool(mcp.NewTool("list_tag_values",
		mcp.WithDescription("List the possible values for a given tag key. Use list_tag_keys to find valid keys."),
		req("key", "Tag key whose values to list"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return textResult(grn("vserver", "server", "tag-value", "--key", sarg(getArgs(r), "key"))), nil
	})

	s.AddTool(mcp.NewTool("create_server_image",
		mcp.WithDescription("Create a user image (custom image) from a vServer instance. Optionally attach tags."),
		req("serverId", "Server ID to create the image from"),
		req("name", "Name of the new image"),
		mcp.WithString("tags", mcp.Description("Optional tags as a comma-separated list of key=value pairs, e.g. \"env=prod,team=infra\"")),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		a := getArgs(r)
		args := []string{"vserver", "server", "create-image",
			"--server-id", sarg(a, "serverId"),
			"--name", sarg(a, "name"),
		}
		if tags := sarg(a, "tags"); tags != "" {
			for _, t := range strings.Split(tags, ",") {
				if t = strings.TrimSpace(t); t != "" {
					args = append(args, "--tag", t)
				}
			}
		}
		return textResult(grn(args...)), nil
	})
}

// ── volume tools ───────────────────────────────────────────────────────────────

func registerVolumeTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool("list_volumes",
		mcp.WithDescription("List all block storage volumes in the project."),
		opt("name"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := []string{"vserver", "volume", "list"}
		if n, ok := getArgs(r)["name"].(string); ok && n != "" {
			args = append(args, "--name", n)
		}
		return textResult(grn(args...)), nil
	})

	s.AddTool(mcp.NewTool("get_volume",
		mcp.WithDescription("Get details of a specific volume."),
		req("volumeId", "Volume ID"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return textResult(grn("vserver", "volume", "get",
			"--volume-id", getArgs(r)["volumeId"].(string),
		)), nil
	})

	s.AddTool(mcp.NewTool("create_volume",
		mcp.WithDescription("Create a new block storage volume. Use list_volume_types to find a valid volumeTypeId."),
		req("name", "Volume name"),
		req("volumeTypeId", "Volume type ID (from list_volume_types)"),
		req("zoneId", "Availability zone ID"),
		mcp.WithNumber("size", mcp.Required(), mcp.Description("Volume size in GiB (minimum 1)")),
		opt("description"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := []string{"vserver", "volume", "create",
			"--name", getArgs(r)["name"].(string),
			"--volume-type-id", getArgs(r)["volumeTypeId"].(string),
			"--zone-id", getArgs(r)["zoneId"].(string),
			"--size", fmt.Sprintf("%d", int(getArgs(r)["size"].(float64))),
		}
		if v, ok := getArgs(r)["description"].(string); ok && v != "" {
			args = append(args, "--description", v)
		}
		return textResult(grn(args...)), nil
	})

	s.AddTool(mcp.NewTool("resize_volume",
		mcp.WithDescription("Resize a volume's size in GiB, change its volume type, or both. At least one of size or newVolumeTypeId must be provided. Use list_volume_types to find a valid newVolumeTypeId. Volume size can only be increased, not decreased."),
		req("volumeId", "Volume ID"),
		mcp.WithNumber("size", mcp.Description("New size in GiB (must be >= current size; omit to keep current size)")),
		opt("newVolumeTypeId"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		a := getArgs(r)
		args := []string{"vserver", "volume", "resize",
			"--volume-id", a["volumeId"].(string),
		}
		if sz, ok := a["size"].(float64); ok && sz > 0 {
			args = append(args, "--size", fmt.Sprintf("%d", int(sz)))
		}
		if v, ok := a["newVolumeTypeId"].(string); ok && v != "" {
			args = append(args, "--volume-type-id", v)
		}
		return textResult(grn(args...)), nil
	})

	s.AddTool(mcp.NewTool("delete_volume",
		mcp.WithDescription("Delete a volume. This action is irreversible. Always confirm with the user before calling this."),
		req("volumeId", "Volume ID"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return textResult(grn("vserver", "volume", "delete", "--force",
			"--volume-id", getArgs(r)["volumeId"].(string),
		)), nil
	})
}

// ── network tools ─────────────────────────────────────────────────────────────

func registerNetworkTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool("list_vpcs",
		mcp.WithDescription("List all VPCs (Virtual Private Clouds) in the project."),
		opt("name"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := []string{"vserver", "vpc", "list"}
		if n, ok := getArgs(r)["name"].(string); ok && n != "" {
			args = append(args, "--name", n)
		}
		return textResult(grn(args...)), nil
	})

	s.AddTool(mcp.NewTool("create_vpc",
		mcp.WithDescription("Create a new VPC with a CIDR block (e.g. 10.0.0.0/16)."),
		req("name", "VPC name"),
		req("cidr", "CIDR block, e.g. 10.0.0.0/16"),
		opt("description"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := []string{"vserver", "vpc", "create",
			"--name", getArgs(r)["name"].(string),
			"--cidr", getArgs(r)["cidr"].(string),
		}
		if v, ok := getArgs(r)["description"].(string); ok && v != "" {
			args = append(args, "--description", v)
		}
		return textResult(grn(args...)), nil
	})

	s.AddTool(mcp.NewTool("list_subnets",
		mcp.WithDescription("List all subnets within a VPC."),
		req("vpcId", "VPC ID"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return textResult(grn("vserver", "subnet", "list",
			"--vpc-id", getArgs(r)["vpcId"].(string),
		)), nil
	})

	s.AddTool(mcp.NewTool("create_subnet",
		mcp.WithDescription("Create a subnet inside a VPC. The CIDR must be within the VPC CIDR range."),
		req("vpcId", "VPC ID"),
		req("cidr", "Subnet CIDR, e.g. 10.0.1.0/24"),
		req("zoneId", "Availability zone ID"),
		opt("name"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := []string{"vserver", "subnet", "create",
			"--vpc-id", getArgs(r)["vpcId"].(string),
			"--cidr", getArgs(r)["cidr"].(string),
			"--zone-id", getArgs(r)["zoneId"].(string),
		}
		if v, ok := getArgs(r)["name"].(string); ok && v != "" {
			args = append(args, "--name", v)
		}
		return textResult(grn(args...)), nil
	})
}

// ── ssh key tools ─────────────────────────────────────────────────────────────

func registerSSHKeyTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool("list_sshkeys",
		mcp.WithDescription("List all SSH keys in the project."),
		opt("name"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := []string{"vserver", "sshkey", "list"}
		if n := sarg(getArgs(r), "name"); n != "" {
			args = append(args, "--name", n)
		}
		return textResult(grn(args...)), nil
	})

	s.AddTool(mcp.NewTool("create_sshkey",
		mcp.WithDescription("Create a new SSH key pair. The server generates both keys; the private key is saved as a '<name>.pem' file (in the Downloads directory by default, or outputDir if given) and is returned only once."),
		req("name", "SSH key name"),
		opt("outputDir"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		a := getArgs(r)
		args := []string{"vserver", "sshkey", "create", "--name", sarg(a, "name")}
		if v := sarg(a, "outputDir"); v != "" {
			args = append(args, "--output-dir", v)
		}
		return textResult(grn(args...)), nil
	})

	s.AddTool(mcp.NewTool("import_sshkey",
		mcp.WithDescription("Import an existing SSH public key. Provide the key contents directly with publicKey, or a path to a public key file with publicKeyFile (exactly one)."),
		req("name", "SSH key name"),
		opt("publicKey"),
		opt("publicKeyFile"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		a := getArgs(r)
		args := []string{"vserver", "sshkey", "import", "--name", sarg(a, "name")}
		if v := sarg(a, "publicKey"); v != "" {
			args = append(args, "--public-key", v)
		}
		if v := sarg(a, "publicKeyFile"); v != "" {
			args = append(args, "--public-key-file", v)
		}
		return textResult(grn(args...)), nil
	})

	s.AddTool(mcp.NewTool("delete_sshkey",
		mcp.WithDescription("Delete an SSH key. This action is irreversible. Always confirm with the user before calling this."),
		req("sshKeyId", "SSH key ID"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return textResult(grn("vserver", "sshkey", "delete", "--force",
			"--sshkey-id", sarg(getArgs(r), "sshKeyId"),
		)), nil
	})
}

// ── placement group tools ─────────────────────────────────────────────────────

func registerPlacementGroupTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool("list_placement_groups",
		mcp.WithDescription("List all placement groups (server groups) in the project."),
		opt("name"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := []string{"vserver", "placement-group", "list"}
		if n := sarg(getArgs(r), "name"); n != "" {
			args = append(args, "--name", n)
		}
		return textResult(grn(args...)), nil
	})

	s.AddTool(mcp.NewTool("list_placement_group_policies",
		mcp.WithDescription("List the placement group policies available for create_placement_group. Returns policy IDs and names. Pass language 'vi' for Vietnamese descriptions (default English)."),
		opt("language"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := []string{"vserver", "placement-group", "list-policies"}
		if v := sarg(getArgs(r), "language"); v != "" {
			args = append(args, "--language", v)
		}
		return textResult(grn(args...)), nil
	})

	s.AddTool(mcp.NewTool("create_placement_group",
		mcp.WithDescription("Create a new placement group. Use list_placement_group_policies first to find a valid policyId."),
		req("name", "Placement group name"),
		req("policyId", "Policy ID (from list_placement_group_policies)"),
		opt("description"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		a := getArgs(r)
		args := []string{"vserver", "placement-group", "create",
			"--name", sarg(a, "name"),
			"--policy-id", sarg(a, "policyId"),
		}
		if v := sarg(a, "description"); v != "" {
			args = append(args, "--description", v)
		}
		return textResult(grn(args...)), nil
	})

	s.AddTool(mcp.NewTool("edit_placement_group",
		mcp.WithDescription("Update a placement group's name and/or description. Only the provided fields are changed."),
		req("placementGroupId", "Placement group (server group) ID"),
		opt("name"),
		opt("description"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		a := getArgs(r)
		args := []string{"vserver", "placement-group", "edit", "--placement-group-id", sarg(a, "placementGroupId")}
		if v := sarg(a, "name"); v != "" {
			args = append(args, "--name", v)
		}
		if v := sarg(a, "description"); v != "" {
			args = append(args, "--description", v)
		}
		return textResult(grn(args...)), nil
	})

	s.AddTool(mcp.NewTool("delete_placement_group",
		mcp.WithDescription("Delete a placement group. This action is irreversible. Always confirm with the user before calling this."),
		req("placementGroupId", "Placement group (server group) ID"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return textResult(grn("vserver", "placement-group", "delete", "--force",
			"--placement-group-id", sarg(getArgs(r), "placementGroupId"),
		)), nil
	})
}

// ── user image tools ──────────────────────────────────────────────────────────

func registerUserImageTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool("list_user_images",
		mcp.WithDescription("List all user images (custom images created from servers) in the project."),
		opt("name"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := []string{"vserver", "user-image", "list"}
		if n := sarg(getArgs(r), "name"); n != "" {
			args = append(args, "--name", n)
		}
		return textResult(grn(args...)), nil
	})

	s.AddTool(mcp.NewTool("update_user_image_tags",
		mcp.WithDescription("Replace the key/value tag list of a user image. The provided tags become the complete set — any tag not included is removed. Tags that were changed should go in editedTags so they are marked isEdited=true."),
		req("userImageId", "User image ID whose tags to update"),
		mcp.WithString("tags", mcp.Description("Unchanged tags as a comma-separated list of key=value pairs (isEdited=false), e.g. \"env=prod,team=infra\"")),
		mcp.WithString("editedTags", mcp.Description("Changed tags as a comma-separated list of key=value pairs (isEdited=true)")),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		a := getArgs(r)
		args := []string{"vserver", "user-image", "update-tags",
			"--user-image-id", sarg(a, "userImageId"),
		}
		appendTagFlags := func(value, flag string) {
			if value == "" {
				return
			}
			for _, t := range strings.Split(value, ",") {
				if t = strings.TrimSpace(t); t != "" {
					args = append(args, flag, t)
				}
			}
		}
		appendTagFlags(sarg(a, "tags"), "--tag")
		appendTagFlags(sarg(a, "editedTags"), "--edited-tag")
		return textResult(grn(args...)), nil
	})

	s.AddTool(mcp.NewTool("delete_user_image",
		mcp.WithDescription("Delete a user image. This action is irreversible. Always confirm with the user before calling this."),
		req("userImageId", "User image ID"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return textResult(grn("vserver", "user-image", "delete", "--force",
			"--user-image-id", sarg(getArgs(r), "userImageId"),
		)), nil
	})
}

// ── floating IP tools ─────────────────────────────────────────────────────────

func registerFloatingIPTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool("list_floating_ips",
		mcp.WithDescription("List all floating IPs (public WAN IP addresses) in the project."),
		opt("name"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := []string{"vserver", "floating-ip", "list"}
		if n := sarg(getArgs(r), "name"); n != "" {
			args = append(args, "--name", n)
		}
		return textResult(grn(args...)), nil
	})

	s.AddTool(mcp.NewTool("delete_floating_ip",
		mcp.WithDescription("Delete a floating IP. This action is irreversible. Always confirm with the user before calling this."),
		req("floatingIpId", "Floating IP ID"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return textResult(grn("vserver", "floating-ip", "delete", "--force",
			"--floating-ip-id", sarg(getArgs(r), "floatingIpId"),
		)), nil
	})
}

// ── network interface tools ───────────────────────────────────────────────────

func registerNetworkInterfaceTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool("list_network_interfaces",
		mcp.WithDescription("List all elastic network interfaces in the project."),
		opt("name"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := []string{"vserver", "network-interface", "list"}
		if n := sarg(getArgs(r), "name"); n != "" {
			args = append(args, "--name", n)
		}
		return textResult(grn(args...)), nil
	})

	s.AddTool(mcp.NewTool("create_network_interface",
		mcp.WithDescription("Create a new elastic network interface in a zone. Optionally attach tags."),
		req("name", "Name of the network interface"),
		req("zoneId", "Availability zone ID (e.g. HCM03-1C)"),
		mcp.WithString("tags", mcp.Description("Optional tags as a comma-separated list of key=value pairs, e.g. \"env=prod,team=infra\"")),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		a := getArgs(r)
		args := []string{"vserver", "network-interface", "create",
			"--name", sarg(a, "name"),
			"--zone-id", sarg(a, "zoneId"),
		}
		if tags := sarg(a, "tags"); tags != "" {
			for _, t := range strings.Split(tags, ",") {
				if t = strings.TrimSpace(t); t != "" {
					args = append(args, "--tag", t)
				}
			}
		}
		return textResult(grn(args...)), nil
	})

	s.AddTool(mcp.NewTool("rename_network_interface",
		mcp.WithDescription("Rename an elastic network interface. Only the name can be changed."),
		req("networkInterfaceId", "Network interface ID"),
		req("name", "New name"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		a := getArgs(r)
		return textResult(grn("vserver", "network-interface", "edit",
			"--network-interface-id", sarg(a, "networkInterfaceId"),
			"--name", sarg(a, "name"),
		)), nil
	})

	s.AddTool(mcp.NewTool("update_network_interface_tags",
		mcp.WithDescription("Replace the key/value tag list of a network interface. The provided tags become the complete set — any tag not included is removed. Tags that were changed should go in editedTags so they are marked isEdited=true."),
		req("networkInterfaceId", "Network interface ID whose tags to update"),
		mcp.WithString("tags", mcp.Description("Unchanged tags as a comma-separated list of key=value pairs (isEdited=false), e.g. \"env=prod,team=infra\"")),
		mcp.WithString("editedTags", mcp.Description("Changed tags as a comma-separated list of key=value pairs (isEdited=true)")),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		a := getArgs(r)
		args := []string{"vserver", "network-interface", "update-tags",
			"--network-interface-id", sarg(a, "networkInterfaceId"),
		}
		appendTagFlags := func(value, flag string) {
			if value == "" {
				return
			}
			for _, t := range strings.Split(value, ",") {
				if t = strings.TrimSpace(t); t != "" {
					args = append(args, flag, t)
				}
			}
		}
		appendTagFlags(sarg(a, "tags"), "--tag")
		appendTagFlags(sarg(a, "editedTags"), "--edited-tag")
		return textResult(grn(args...)), nil
	})

	s.AddTool(mcp.NewTool("delete_network_interface",
		mcp.WithDescription("Delete an elastic network interface. This action is irreversible. Always confirm with the user before calling this."),
		req("networkInterfaceId", "Network interface ID"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return textResult(grn("vserver", "network-interface", "delete", "--force",
			"--network-interface-id", sarg(getArgs(r), "networkInterfaceId"),
		)), nil
	})
}

// ── DHCP option tools ─────────────────────────────────────────────────────────

func registerDhcpTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool("list_dhcp_options",
		mcp.WithDescription("List all DHCP option sets in the project."),
		opt("name"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := []string{"vserver", "dhcp", "list"}
		if n := sarg(getArgs(r), "name"); n != "" {
			args = append(args, "--name", n)
		}
		return textResult(grn(args...)), nil
	})

	s.AddTool(mcp.NewTool("create_dhcp_option",
		mcp.WithDescription("Create a new DHCP option set. The two default DNS servers (10.166.12.196, 10.166.12.197) are always included; up to 2 more may be added via dnsServers (4 total)."),
		req("name", "Name of the DHCP option set"),
		mcp.WithString("dnsServers", mcp.Description("Additional DNS server IPs as a comma-separated list (max 2), e.g. \"10.0.0.1,10.0.0.2\"")),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		a := getArgs(r)
		args := []string{"vserver", "dhcp", "create", "--name", sarg(a, "name")}
		if dns := sarg(a, "dnsServers"); dns != "" {
			for _, ip := range strings.Split(dns, ",") {
				if ip = strings.TrimSpace(ip); ip != "" {
					args = append(args, "--dns-server", ip)
				}
			}
		}
		return textResult(grn(args...)), nil
	})

	s.AddTool(mcp.NewTool("get_dhcp_option",
		mcp.WithDescription("Get the full details of a single DHCP option set by its ID."),
		req("dhcpOptionId", "DHCP option ID"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return textResult(grn("vserver", "dhcp", "get",
			"--dhcp-option-id", sarg(getArgs(r), "dhcpOptionId"),
		)), nil
	})

	s.AddTool(mcp.NewTool("list_dhcp_option_vpcs",
		mcp.WithDescription("List the VPCs (networks) currently associated with a DHCP option set."),
		req("dhcpOptionId", "DHCP option ID"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return textResult(grn("vserver", "dhcp", "list-vpcs",
			"--dhcp-option-id", sarg(getArgs(r), "dhcpOptionId"),
		)), nil
	})

	s.AddTool(mcp.NewTool("associate_vpc_dhcp_option",
		mcp.WithDescription("Associate a VPC with a DHCP option set, or detach it. A VPC can belong to only one DHCP option set. Provide dhcpOptionId to associate; set detach=true to remove the VPC's association."),
		req("vpcId", "VPC (network) ID to update"),
		opt("dhcpOptionId"),
		mcp.WithBoolean("detach", mcp.Description("Detach the VPC from its current DHCP option set (default false)")),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		a := getArgs(r)
		args := []string{"vserver", "dhcp", "associate-vpc", "--vpc-id", sarg(a, "vpcId")}
		if d, _ := a["detach"].(bool); d {
			args = append(args, "--detach")
		} else if id := sarg(a, "dhcpOptionId"); id != "" {
			args = append(args, "--dhcp-option-id", id)
		}
		return textResult(grn(args...)), nil
	})

	s.AddTool(mcp.NewTool("delete_dhcp_option",
		mcp.WithDescription("Delete a DHCP option set. This action is irreversible. Always confirm with the user before calling this."),
		req("dhcpOptionId", "DHCP option ID"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return textResult(grn("vserver", "dhcp", "delete", "--force",
			"--dhcp-option-id", sarg(getArgs(r), "dhcpOptionId"),
		)), nil
	})
}

// ── security group tools ──────────────────────────────────────────────────────

func registerSecgroupTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool("list_secgroups",
		mcp.WithDescription("List all security groups in the project."),
		opt("name"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := []string{"vserver", "secgroup", "list"}
		if n, ok := getArgs(r)["name"].(string); ok && n != "" {
			args = append(args, "--name", n)
		}
		return textResult(grn(args...)), nil
	})

	s.AddTool(mcp.NewTool("create_secgroup",
		mcp.WithDescription("Create a new security group."),
		req("name", "Security group name"),
		opt("description"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := []string{"vserver", "secgroup", "create",
			"--name", getArgs(r)["name"].(string),
		}
		if v, ok := getArgs(r)["description"].(string); ok && v != "" {
			args = append(args, "--description", v)
		}
		return textResult(grn(args...)), nil
	})

	s.AddTool(mcp.NewTool("list_secgroup_rules",
		mcp.WithDescription("List all rules in a security group."),
		req("secgroupId", "Security group ID"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return textResult(grn("vserver", "secgroup", "rule", "list",
			"--secgroup-id", getArgs(r)["secgroupId"].(string),
		)), nil
	})

	s.AddTool(mcp.NewTool("create_secgroup_rule",
		mcp.WithDescription("Add a rule to a security group. Protocol 'icmp' and 'any' do not use port ranges."),
		req("secgroupId", "Security group ID"),
		req("direction", "Traffic direction: ingress or egress"),
		req("protocol", "Protocol: tcp, udp, icmp, or any"),
		req("remoteIpPrefix", "Remote CIDR, e.g. 0.0.0.0/0"),
		req("etherType", "Ether type: IPv4 or IPv6"),
		mcp.WithNumber("portRangeMin", mcp.Description("Min port (tcp/udp only)")),
		mcp.WithNumber("portRangeMax", mcp.Description("Max port (tcp/udp only)")),
		opt("description"),
	), func(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := []string{"vserver", "secgroup", "rule", "create",
			"--secgroup-id", getArgs(r)["secgroupId"].(string),
			"--direction", getArgs(r)["direction"].(string),
			"--protocol", getArgs(r)["protocol"].(string),
			"--remote-ip-prefix", getArgs(r)["remoteIpPrefix"].(string),
			"--ether-type", getArgs(r)["etherType"].(string),
		}
		if v, ok := getArgs(r)["portRangeMin"].(float64); ok {
			args = append(args, "--port-range-min", fmt.Sprintf("%d", int(v)))
		}
		if v, ok := getArgs(r)["portRangeMax"].(float64); ok {
			args = append(args, "--port-range-max", fmt.Sprintf("%d", int(v)))
		}
		if v, ok := getArgs(r)["description"].(string); ok && v != "" {
			args = append(args, "--description", v)
		}
		return textResult(grn(args...)), nil
	})
}
