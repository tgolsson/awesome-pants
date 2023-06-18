package website

import (
	"sort"

	"github.com/tgolsson/awesome-pants/src/go/plugins"
	"github.com/tgolsson/awesome-pants/src/go/recipes"
)

// Generate generates the website.
func Generate(plugins *plugins.PluginsCache, recipes *recipes.AdhocRecipes) (map[string][]byte, error) {
	sort.Slice(plugins.Plugins, func(i, j int) bool {
		return plugins.Plugins[i].Name < plugins.Plugins[j].Name
	})

	output := make(map[string][]byte)

	html, err := genHomeHtml()
	if err != nil {
		return nil, err
	}

	output["index.html"] = []byte(html)

	html, err = genPluginHtml(plugins)
	if err != nil {
		return nil, err
	}

	output["plugins/index.html"] = []byte(html)

	html, err = genRecipesHtml(recipes)
	if err != nil {
		return nil, err
	}

	output["recipes/index.html"] = []byte(html)
	return output, nil
}
