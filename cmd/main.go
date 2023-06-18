package main

import (
	"fmt"
	"strings"

	"github.com/alecthomas/kong"
)

var CLI struct {
	Refresh struct {
		Path   string `arg:"" name:"path" help:"Paths to refresh." type:"path"`
		Output string `arg:"" name:"output" help:"Output path." type:"path"`
	} `cmd:"" help:"Refreshes the plugins."`

	Gen struct {
		Path        string `arg:"" name:"path" help:"Paths to generate." type:"path"`
		RecipesPath string `arg:"" name:"recipes-path" help:"Path to recipes." type:"path"`
		Output      string `arg:"" name:"output" help:"Output path." type:"path"`
	} `cmd:"" help:"Generates the plugins."`

	GenReadme struct {
		Path         string `arg:"" name:"path" help:"Paths to generate." type:"path"`
		RecipesPath  string `arg:"" name:"recipes-path" help:"Path to recipes." type:"path"`
		TemplatePath string `arg:"" name:"template-path" help:"Path to template." type:"path"`
		Output       string `arg:"" name:"output" help:"Output path." type:"path"`
	} `cmd:"" help:"Generates the plugins."`
}

func main() {
	ctx := kong.Parse(&CLI)

	switch strings.TrimSpace(ctx.Command()) {
	case "refresh <path> <output>":
		refresh(CLI.Refresh.Path, CLI.Refresh.Output)
	case "gen <path> <recipes-path> <output>":
		err := gen(CLI.Gen.Path, CLI.Gen.RecipesPath, CLI.Gen.Output)
		if err != nil {
			fmt.Println(err)
		}
	case "gen-readme <path> <recipes-path> <template-path> <output>":
		err := genReadme(CLI.GenReadme.Path, CLI.GenReadme.RecipesPath, CLI.GenReadme.TemplatePath, CLI.GenReadme.Output)
		if err != nil {
			fmt.Println(err)
		}

	default:
		fmt.Println("Unknown command: '", ctx.Command(), "'")
		ctx.PrintUsage(true)
	}
}
