package website

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"strings"

	"github.com/tgolsson/awesome-pants/src/go/plugins"
)

func genMain(sb *strings.Builder, cache_ interface{}) error {
	cache := cache_.(*plugins.PluginsCache)
	sb.WriteString("<h2>Plugins</h2>")
	sb.WriteString("<p>Plugins are Python packages that can be used to extend the functionality of Pants. To add plugins to your project, add them to your <code>pants.toml</code> file.</p>")
	sb.WriteString("<p>First, you need to add the plugins as dependencies:")
	sb.WriteString(`<pre><code>[GLOBAL]
plugins = [
  "name_of_pypi_package==1.0.0",
]</code></pre>`)

	sb.WriteString("<p>Then, you may need to add the backends to the <code>backend_packages</code> list</p> if it's not enabled by default:")
	sb.WriteString(`<pre><code>[GLOBAL]
backend_packages = [
  // ... other backends
  "name_of_pypi_package.backend",
]</code></pre>`)

	for _, plugin := range cache.Plugins {
		topHeaderList := ""
		githubIcon := "https://icongr.am/simple/github.svg?size=24&colored=true"
		pypiIcon := "https://icongr.am/simple/pypi.svg?size=24&colored=true"
		hashIcon := "https://icongr.am/feather/hash.svg?size=24&color=currentColor"
		topHeaderList += fmt.Sprintf("<a class=\"icon\" href=\"%s\"><img src=\"%s\" alt=\"Icon\" /></a>", plugin.Pypi, pypiIcon)
		topHeaderList += fmt.Sprintf("<a class=\"icon\" href=\"%s\"><img src=\"%s\" alt=\"Github\" /></a>", plugin.Metadata.Website, githubIcon)
		topHeaderList += fmt.Sprintf("<a class=\"icon\" href=\"#%s\"><img src=\"%s\" alt=\"Permalink\" /></a>", plugin.Name, hashIcon)
		iconList := fmt.Sprintf("<img src=\"%s&size=32\" alt=\"%s\" />", plugin.Icon, plugin.Name)

		badges := fmt.Sprintf(`
			<img alt="Version" src="https://img.shields.io/badge/PyPi-%s-fe7d37">
            <img alt="Recent downloads" src="https://img.shields.io/badge/Recent%%20downloads-%v-green">
`, plugin.Metadata.LatestVersion, humanize.SI(float64(plugin.Stats.Month), ""))

		if plugin.Metadata.License != "" {
			badges += fmt.Sprintf(`
			<img alt="Version" src="https://img.shields.io/badge/License-%s-blue">
`, plugin.Metadata.License)
		}

		if plugin.Status != "" {
			color := "blue"
			if plugin.Status == "Stable" {
				color = "green"
			} else if plugin.Status == "Deprecated" {
				color = "red"
			} else if plugin.Status == "Experimental" {
				color = "yellow"
			}

			badges += fmt.Sprintf(`
			<img alt="Version" src="https://img.shields.io/badge/Status-%s-%s">
`, plugin.Status, color)
		}
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
			</div>`, topHeaderList, plugin.Pypi, plugin.Name, plugin.Metadata.Summary, iconList, badges))

	}

	return nil
}

func genBody(sb *strings.Builder, cache interface{}) error {
	return genScopedWithClass(sb, "div", "container", cache, func(sb *strings.Builder, cache interface{}) error {
		if err := genScoped(sb, "header", cache, func(sb *strings.Builder, cache interface{}) error {
			return genHeader(sb, cache, "..")
		}); err != nil {
			return err
		}
		return genScoped(sb, "main", cache, genMain)
	})
}

// Generates an index.html file which lists all plugins
// and their respective documentation.
func genPluginHtml(cache *plugins.PluginsCache) (string, error) {
	var sb strings.Builder
	err := genScoped(&sb, "html", cache, func(sb *strings.Builder, cache interface{}) error {
		err := genScoped(sb, "head", cache, genHead)
		if err != nil {
			return err
		}
		return genScoped(sb, "body", cache, genBody)
	})
	if err != nil {
		return "", err
	}
	return sb.String(), nil
}
