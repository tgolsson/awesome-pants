package readme

import (
	"fmt"
	"strings"

	"github.com/tgolsson/awesome-pants/src/go/recipes"
)

func generateRecipes(recipes *recipes.AdhocRecipes) string {
	var sb strings.Builder
	sb.WriteString("## Recipes\n\n")
	for _, recipe := range recipes.Recipes {
		if recipe.AuthorGithub == "" {
			sb.WriteString(fmt.Sprintf("- [%s](%s) by %s  \n", recipe.Name, recipe.Url, recipe.Author))
		} else {
			sb.WriteString(fmt.Sprintf("- [%s](%s) by [%s](https://github.com/%s)  \n", recipe.Name, recipe.Url, recipe.Author, recipe.AuthorGithub))
		}
		sb.WriteString(fmt.Sprintf("  %s\n\n", recipe.Summary))
	}
	sb.WriteString("\n")
	return sb.String()

}
