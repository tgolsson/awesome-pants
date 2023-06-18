package recipes

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

// AdhocRecipe is a single recipe item.
type AdhocRecipe struct {
	Name         string   `toml:"name"`
	Summary      string   `toml:"summary"`
	Author       string   `toml:"author"`
	AuthorGithub string   `toml:"author_github"`
	Url          string   `toml:"url"`
	Icons        []string `toml:"icons"`
}

// AdhocRecipes is a collection of recipes.
type AdhocRecipes struct {
	Recipes []AdhocRecipe `toml:"recipes"`
}

// LoadRecipes loads recipes from a TOML file.
func LoadRecipes(path string) (*AdhocRecipes, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data AdhocRecipes
	if _, err := toml.Decode(string(content), &data); err != nil {
		return nil, err
	}

	return &data, nil
}
