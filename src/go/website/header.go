package website

import (
	"strings"
)

func genHeader(sb *strings.Builder, cache interface{}) error {
	genScopedWithClass(sb, "nav", "nav", cache, func(sb *strings.Builder, cache interface{}) error {
		if err := genScopedWithClass(sb, "div", "nav-left", cache, func(sb *strings.Builder, cache interface{}) error {
			sb.WriteString("<a class=\"/\" href=\"/\">Home</a>")
			sb.WriteString("<a href=\"/plugins\">Plugins</a>")
			sb.WriteString("<a href=\"/recipes\">Adhoc Recipes</a>")
			return nil
		}); err != nil {
			return err
		}
		if err := genScopedWithClass(sb, "div", "nav-center", cache, func(sb *strings.Builder, cache interface{}) error {
			sb.WriteString("<a class=\"brand\" href=\"/\">Awesome Pants</a>")
			return nil
		}); err != nil {
			return err
		}

		if err := genScopedWithClass(sb, "div", "nav-right", cache, func(sb *strings.Builder, cache interface{}) error {
			sb.WriteString("<a href=\"https://github.com/tgolsson/awesome-pants\">Contribute</a>")
			return nil
		}); err != nil {
			return err
		}

		return nil
	})
	sb.WriteString("<h6 class=\"is-horizontal-align\">Plugins and tools for the Pants Ecosystem</h6>")
	sb.WriteString("<hr />")
	return nil
}
