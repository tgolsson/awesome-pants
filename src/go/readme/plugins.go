package readme

import (
	"fmt"
	"strings"

	"github.com/tgolsson/awesome-pants/src/go/plugins"
)

func generatePlugins(plugins *plugins.PluginsCache) string {
	var sb strings.Builder
	sb.WriteString("## Plugins\n\n")
	for _, plugin := range plugins.Plugins {
		sb.WriteString(fmt.Sprintf("- [%s](%s) - Version: %s  ", plugin.Name, plugin.Pypi, plugin.Metadata.LatestVersion))
		sb.WriteString(fmt.Sprintf("  %s\n\n", plugin.Metadata.Summary))
	}
	sb.WriteString("\n")
	return sb.String()
}
