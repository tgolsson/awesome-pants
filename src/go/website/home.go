package website

import (
	"strings"
)

func genHomeBody(sb *strings.Builder, cache interface{}) error {
	return genScopedWithClass(sb, "div", "container", cache, func(sb *strings.Builder, cache interface{}) error {
		if err := genScoped(sb, "header", cache, func(sb *strings.Builder, cache interface{}) error {
			return genHeader(sb, cache, ".")
		}); err != nil {
			return err
		}
		return genScoped(sb, "main", cache, func(sb *strings.Builder, cache interface{}) error {
			intro := `
				<p>
					Awesome Pants is a landing page to collect all the awesome plugins for <a href="https://pantsbuild.org">Pants</a>, and resources for plugin development.
				</p>
				<p>
If you have a plugin you'd like to add, please <a href="https://github.com/tgolsson-awesome-pants">open a PR</a>! If you're looking for a plugin, check out the <a href="/plugins">plugins page</a>.
				</p>
			`
			sb.WriteString(intro)
			return nil
		})
	})
}

func genHomeHtml() (string, error) {
	var sb strings.Builder
	err := genScoped(&sb, "html", nil, func(sb *strings.Builder, cache interface{}) error {
		err := genScoped(sb, "head", cache, genHead)
		if err != nil {
			return err
		}
		return genScoped(sb, "body", cache, genHomeBody)
	})
	if err != nil {
		return "", err
	}
	return sb.String(), nil
}
