# GreenNode CLI

[![CI](https://github.com/fuochuy/grn-vserver-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/fuochuy/grn-vserver-cli/actions/workflows/ci.yml)
[![Release](https://img.shields.io/github/v/release/vngcloud/greennode-cli)](https://github.com/fuochuy/grn-vserver-cli/releases/latest)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![Go](https://img.shields.io/badge/Go-1.25+-00ADD8.svg)](https://go.dev)

Universal Command Line Interface for **GreenNode** (powered by VNG Cloud).

`grn-vserver` lets you manage GreenNode services  server, volume, and more — directly from the terminal. Written in Go, distributed as a single static binary with zero runtime dependencies.

---

- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [MCP Integration](#mcp-integration)
- [Contributing](#contributing)
- [Security](#security)
- [License](#license)

---

## Installation

Download the latest binary for your platform from [GitHub Releases](https://github.com/fuochuy/grn-vserver-cli/releases/latest):

| Platform | Architecture | Binary |
|----------|-------------|--------|
| macOS | Apple Silicon (M1/M2/M3) | `grn-darwin-arm64` |
| macOS | Intel | `grn-darwin-amd64` |
| Linux | x86_64 | `grn-linux-amd64` |
| Linux | ARM64 | `grn-linux-arm64` |
| Windows | x86_64 (64-bit) | `grn-windows-amd64.exe` |
| Windows | x86 (32-bit) | `grn-windows-386.exe` |

```bash
# macOS Apple Silicon
curl -L -o grn https://github.com/fuochuy/grn-vserver-cli/releases/latest/download/grn-darwin-arm64
chmod +x grn
sudo mv grn /usr/local/bin/

# Verify
grn --version
# grn-cli/1.5.0 Go/1.25.x darwin/arm64
```

**One-line installer** (macOS / Linux):

```bash
curl -fsSL https://raw.githubusercontent.com/vngcloud/greennode-cli/main/scripts/install | bash
```

**Build from source** (requires [Go 1.25+](https://go.dev/dl/)):

```bash
git clone https://github.com/fuochuy/grn-vserver-cli.git
cd greennode-cli/go
go build -o grn .
sudo mv grn /usr/local/bin/
```

Full installation guide: [docs/installation.md](docs/installation.md)

---

## Configuration

```bash
grn configure
```

The wizard prompts for Client ID, Client Secret, region, and project ID. Credentials are obtained from the [VNG Cloud IAM Portal](https://hcm-3.console.vngcloud.vn/iam/) under **Service Accounts**.

```
GRN Client ID [None]: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
GRN Client Secret [None]: ****************
Default region name [HCM-3]:
Default output format [json]:
Project ID (leave blank to auto-detect) [None]:
Auto-detected project_id: pro-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
```

Verify saved config:

```bash
grn configure list
```

Full configuration reference (environment variables, profiles, config files): [docs/configuration.md](docs/configuration.md)

---

## Usage

### vServer — virtual machines

```bash
grn vserver server list
grn vserver server create --name my-server --zone-id <z> --image-id <i> --flavor-id <f> ...
grn vserver server start  --server-id <id>
grn vserver server stop   --server-id <id>
grn vserver server delete --server-id <id>

grn vserver volume list
grn vserver secgroup list
grn vserver vpc list
```

### Output & filtering

### Dry-run

```bash
grn vserver server create --dry-run --name test ...
```

Usage guide: [docs/usage/](docs/usage/README.md) · Full command reference: [docs/commands/](docs/commands/)

---

## MCP Integration

`grn` ships a built-in [MCP](https://modelcontextprotocol.io) server that exposes all commands as tools for AI assistants (Claude Desktop, Cursor, etc.):

```bash
grn mcp   # starts MCP server on stdio
```

Add to `~/Library/Application Support/Claude/claude_desktop_config.json` (macOS):

```json
{
  "mcpServers": {
    "greennode": {
      "command": "/usr/local/bin/grn",
      "args": ["mcp"]
    }
  }
}
```

Then restart Claude Desktop and say things like: *"List my vServer instances"* or *"Create a cluster named prod in zone HCM03-1A"*.

---

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for the full workflow.

Quick start:

```bash
git clone https://github.com/fuochuy/grn-vserver-cli.git
cd greennode-cli/go
go build -o grn .
./grn --version
```

Every PR must include a changelog fragment:

---

## Security

To report a security vulnerability, please see [SECURITY.md](SECURITY.md). Do not open a public issue.

---

## License

Apache 2.0 — see [LICENSE](LICENSE).
