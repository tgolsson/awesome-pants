package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/tgolsson/awesome-pants/src/go/plugins"
	"github.com/tgolsson/awesome-pants/src/go/recipes"
	"github.com/tgolsson/awesome-pants/src/go/website"
)

func ensureOutputPath(path string) (string, error) {
	// Has to be a directory
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		// If the path does not exist, create it
		if os.IsNotExist(err) {
			err := os.MkdirAll(path, 0755)
			if err != nil {
				return "", err
			}
			return path, nil
		}
		return "", err
	}

	// If the path exists, it has to be a directory
	if !fileInfo.IsDir() {
		return "", fmt.Errorf("path %s is not a directory", path)
	}

	return path, nil
}

func writeHtml(path, filename string, html []byte) error {
	path, err := ensureOutputPath(path)
	if err != nil {
		return err
	}

	fmt.Printf("Writing %s%s\n", path, filename)
	return ioutil.WriteFile(fmt.Sprintf("%s%s", path, filename), html, 0644)
}

func gen(path, recipesPath, outputPath string) error {
	cache, err := plugins.LoadCollection(path)
	if err != nil {
		return err
	}

	recipes, err := recipes.LoadRecipes(recipesPath)
	if err != nil {
		return err
	}

	files, err := website.Generate(cache, recipes)
	if err != nil {
		return err
	}

	for filename, html := range files {
		fullPath := fmt.Sprintf("%s/%s", outputPath, filename)
		dirName := fullPath[:strings.LastIndex(fullPath, "/")]
		fileName := fullPath[strings.LastIndex(fullPath, "/")+1:]

		err := writeHtml(dirName, fileName, html)
		if err != nil {
			return err
		}
	}

	return nil
}
