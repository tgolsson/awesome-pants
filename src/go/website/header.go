package website

import (
	"fmt"
	"strings"
)

func genHeader(sb *strings.Builder, cache interface{}, linkPrefix string) error {
	genScopedWithClass(sb, "nav", "nav", cache, func(sb *strings.Builder, cache interface{}) error {
		if err := genScopedWithClass(sb, "div", "nav-left", cache, func(sb *strings.Builder, cache interface{}) error {
			sb.WriteString(fmt.Sprintf("<a href=\"%s/\">Home</a>", linkPrefix))
			sb.WriteString(fmt.Sprintf("<a href=\"%s/plugins\">Plugins</a>", linkPrefix))
			sb.WriteString(fmt.Sprintf("<a href=\"%s/recipes\">Recipes</a>", linkPrefix))
			return nil
		}); err != nil {
			return err
		}
		if err := genScopedWithClass(sb, "div", "nav-center", cache, func(sb *strings.Builder, cache interface{}) error {
			sb.WriteString(fmt.Sprintf("<a href=\"%s/\">Awesome Pants</a>", linkPrefix))
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
