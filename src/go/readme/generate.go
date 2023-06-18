package readme

import (
	"sort"

	"github.com/tgolsson/awesome-pants/src/go/plugins"
	"github.com/tgolsson/awesome-pants/src/go/recipes"
)

// Generate generates the website.
func Generate(plugins *plugins.PluginsCache, recipes *recipes.AdhocRecipes) (string, error) {
	sort.Slice(plugins.Plugins, func(i, j int) bool {
		return plugins.Plugins[i].Name < plugins.Plugins[j].Name
	})

	output := generatePlugins(plugins)
	output += generateRecipes(recipes)

	return output, nil
}
