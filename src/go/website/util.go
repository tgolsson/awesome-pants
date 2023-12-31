package website

import (
	"fmt"
	"strings"
)

func genScoped(sb *strings.Builder, scope string, cache interface{}, callback func(*strings.Builder, interface{}) error) error {
	sb.WriteString(fmt.Sprintf("<%s>", scope))
	err := callback(sb, cache)
	if err != nil {
		return err
	}
	sb.WriteString(fmt.Sprintf("</%s>", scope))
	return nil
}

func genScopedWithClass(sb *strings.Builder, scope string, class string, cache interface{}, callback func(*strings.Builder, interface{}) error) error {
	sb.WriteString(fmt.Sprintf("<%s class=\"%s\">", scope, class))
	err := callback(sb, cache)
	if err != nil {
		return err
	}
	sb.WriteString(fmt.Sprintf("</%s>", scope))
	return nil
}

func genHead(sb *strings.Builder, cache interface{}, pageTitle string) error {
	linkPrefix := "."
	if pageTitle == "" {
		pageTitle = "Awesome Pants"
	} else {
		linkPrefix = ".."
		pageTitle = fmt.Sprintf("Awesome Pants - %s", pageTitle)
	}

	sb.WriteString(fmt.Sprintf("<title>%s</title>", pageTitle))
	sb.WriteString(`<link href="
https://cdn.jsdelivr.net/npm/chota@0.9.2/dist/chota.min.css
" rel="stylesheet">`)
	sb.WriteString(`<style>
    body.dark {
      --bg-color: #000;
      --bg-secondary-color: #131316;
      --font-color: #f5f5f5;
      --color-grey: #ccc;
      --color-darkGrey: #777;
    }
    .icon {
 	 margin-left: 0.5rem;
	 margin-right: 0.5rem;
	}
    .dimmed {
       color: var(--color-grey);
    }
    .flexy {
      display: flex;
	  flex-direction: column;
	  justify-content: space-between;
    }

    .hidden {
      visibility: hidden;
    }
  </style>
  <script>
    if (window.matchMedia &&
        window.matchMedia('(prefers-color-scheme: dark)').matches) {
      document.body.classList.add('dark');
    }
  </script>`)
	sb.WriteString(fmt.Sprintf(`
<link rel="apple-touch-icon" sizes="180x180" href="%s/apple-touch-icon.png">
<link rel="icon" type="image/png" sizes="32x32" href="%s/favicon-32x32.png">
<link rel="icon" type="image/png" sizes="16x16" href="%s/favicon-16x16.png">
<link rel="manifest" href="%s/site.webmanifest">
`, linkPrefix, linkPrefix, linkPrefix, linkPrefix))

	return nil
}
