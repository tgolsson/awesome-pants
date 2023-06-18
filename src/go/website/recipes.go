package website

import (
	"fmt"
	"strings"

	"github.com/tgolsson/awesome-pants/src/go/recipes"
)

func genRecipesMain(sb *strings.Builder, cache_ interface{}) error {
	cache := cache_.(*recipes.AdhocRecipes)
	for _, recipe := range cache.Recipes {
		topHeaderList := ""
		userIcon := "https://icongr.am/octicons/person.svg?size=24&color=currentColor"
		hashIcon := "https://icongr.am/feather/hash.svg?size=24&color=currentColor"

		if recipe.AuthorGithub != "" {
			topHeaderList += fmt.Sprintf("<a class=\"icon\" href=\"https://github.com/%s\"><img src=\"%s\" alt=\"Github\" /></a>", recipe.AuthorGithub, userIcon)
		}
		topHeaderList += fmt.Sprintf("<a class=\"icon\" href=\"#%s\"><img src=\"%s\" alt=\"Permalink\" /></a>", recipe.Name, hashIcon)
		iconList := ""
		for _, icon := range recipe.Icons {
			if strings.HasPrefix(icon, "https://icon") {
				iconList += fmt.Sprintf("<img class=\"icon\" src=\"%s&size=32\" alt=\"%s\" />", icon, recipe.Name)
			} else {
				iconList += fmt.Sprintf("<img class=\"icon\" src=\"%s\" alt=\"%s\" height=\"32px\" />", icon, recipe.Name)
			}
		}

		badges := fmt.Sprintf(`
			<img alt="Version" src="https://img.shields.io/badge/Author-%s-green">
`, recipe.Author)
		sb.WriteString(fmt.Sprintf(`
			<div class="card col flexy">
				<div>
					<div style="float: right;" class="is-right">%s</div>
					<header>
						<h4><a href="%s">%s</a></h4>
					</header>
					<div class="card-body">
						<p>%s</p>
					</div>
					%s
				</div>
				<hr />
				<footer >
%s
				</footer>
			</div>`, topHeaderList, recipe.Url, recipe.Name, recipe.Summary, iconList, badges))
	}

	return nil
}
func genRecipesBody(sb *strings.Builder, cache interface{}) error {
	return genScopedWithClass(sb, "div", "container", cache, func(sb *strings.Builder, cache interface{}) error {
		if err := genScoped(sb, "header", cache, func(sb *strings.Builder, cache interface{}) error {
			return genHeader(sb, cache, "..")
		}); err != nil {
			return err
		}
		return genScoped(sb, "main", cache, genRecipesMain)
	})
}

func genRecipesHtml(cache *recipes.AdhocRecipes) (string, error) {
	var sb strings.Builder
	err := genScoped(&sb, "html", cache, func(sb *strings.Builder, cache interface{}) error {
		err := genScoped(sb, "head", cache, genHead)
		if err != nil {
			return err
		}
		return genScoped(sb, "body", cache, genRecipesBody)
	})
	if err != nil {
		return "", err
	}
	return sb.String(), nil
}
