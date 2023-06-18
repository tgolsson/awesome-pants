package main

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/tgolsson/awesome-pants/src/go/plugins"
	"github.com/tgolsson/awesome-pants/src/go/readme"
	"github.com/tgolsson/awesome-pants/src/go/recipes"
)

func readTemplate(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	template, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(template), nil
}

func writeOutput(path, content string) error {
	return ioutil.WriteFile(path, []byte(content), 0644)
}

func genReadme(path, recipesPath, templatePath, outputPath string) error {
	cache, err := plugins.LoadCollection(path)
	if err != nil {
		return err
	}

	recipes, err := recipes.LoadRecipes(recipesPath)
	if err != nil {
		return err
	}

	content, err := readme.Generate(cache, recipes)
	if err != nil {
		return err
	}

	template, err := readTemplate(templatePath)
	if err != nil {
		return err
	}

	output := strings.Replace(template, "{{CONTENT}}", content, 1)

	return writeOutput(outputPath, output)
}
