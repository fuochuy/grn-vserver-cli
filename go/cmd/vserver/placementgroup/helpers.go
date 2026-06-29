package placementgroup

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vngcloud/greennode-cli/internal/client"
	"github.com/vngcloud/greennode-cli/internal/config"
	"github.com/vngcloud/greennode-cli/internal/vserverclient"
)

func createClient(cmd *cobra.Command) (*client.GreenNodeClient, *config.Config, error) {
	return vserverclient.BuildClient(cmd)
}

func getProjectID(cfg *config.Config) (string, error) {
	return vserverclient.ProjectID(cfg)
}

func outputResult(cmd *cobra.Command, cfg *config.Config, data interface{}) error {
	return vserverclient.Output(cmd, cfg, data)
}

// resolveOutput returns the effective output format, mirroring vserverclient.Output:
// the --output flag, falling back to the configured default, then "json".
func resolveOutput(cmd *cobra.Command, cfg *config.Config) string {
	output, _ := cmd.Flags().GetString("output")
	if output == "" && cfg != nil {
		output = cfg.Output
	}
	if output == "" {
		output = "json"
	}
	return output
}

// uuidPreviewLen is how many characters of the uuid are shown in table output.
const uuidPreviewLen = 20

// tableColumns is the column order shown in table output. Fields not listed here
// (e.g. policyId, serverGroupId) are hidden from the table but remain in JSON.
var tableColumns = []string{"uuid", "name", "policyName", "description", "servers", "createdAt"}

// truncateID shortens a long id to a readable preview for table output.
func truncateID(s string) string {
	if len(s) <= uuidPreviewLen {
		return s
	}
	return s[:uuidPreviewLen] + "…"
}

// serverNames turns a "servers" array of {name, uuid} objects into a comma-separated
// list of names, so the table column shows names instead of nested maps.
func serverNames(v interface{}) interface{} {
	items, ok := v.([]interface{})
	if !ok {
		return v
	}
	names := make([]string, 0, len(items))
	for _, it := range items {
		if m, ok := it.(map[string]interface{}); ok {
			if n, ok := m["name"].(string); ok && n != "" {
				names = append(names, n)
				continue
			}
		}
		names = append(names, fmt.Sprint(it))
	}
	return strings.Join(names, ", ")
}

// transformForTable adapts a placement group response for table output: it shortens
// the uuid and renders servers as a comma-separated list of names. It is only applied
// for table output — JSON keeps the full response.
func transformForTable(v interface{}) interface{} {
	switch t := v.(type) {
	case map[string]interface{}:
		out := make(map[string]interface{}, len(t))
		for k, val := range t {
			switch k {
			case "uuid":
				if s, ok := val.(string); ok {
					out[k] = truncateID(s)
					continue
				}
				out[k] = val
			case "servers":
				out[k] = serverNames(val)
			default:
				out[k] = transformForTable(val)
			}
		}
		return out
	case []interface{}:
		out := make([]interface{}, len(t))
		for i, item := range t {
			out[i] = transformForTable(item)
		}
		return out
	default:
		return v
	}
}

// outputGroupList prints a placement group list. For table output it applies
// table-friendly transforms and a fixed column order; other formats (including
// JSON) show the full response.
func outputGroupList(cmd *cobra.Command, cfg *config.Config, result interface{}) error {
	if resolveOutput(cmd, cfg) == "table" {
		return vserverclient.OutputWithColumns(cmd, cfg, transformForTable(result), tableColumns)
	}
	return outputResult(cmd, cfg, result)
}
