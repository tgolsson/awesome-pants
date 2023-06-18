package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/BurntSushi/toml"

	"github.com/tgolsson/awesome-pants/src/go/plugins"
	"github.com/tgolsson/awesome-pants/src/go/recipes"
	"github.com/tgolsson/awesome-pants/src/go/website"
)

type ManifestData struct {
	Static []string `toml:"static"`
}
type Manifest struct {
	ManifestData ManifestData `toml:"manifest"`
}

func loadManifest(path string) (*Manifest, error) {
	var m Manifest
	_, err := toml.DecodeFile(path, &m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func copyFile(src, dst string) error {
	fmt.Printf("Copying %s to %s\n", src, dst)
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	filename := filepath.Base(src)
	destination, err := os.Create(filepath.Join(dst, filename))
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

func copyDir(src, dst string) error {
	fmt.Printf("Copying %s to %s\n", src, dst)
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		sourcePath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			return err
		}

		stat, ok := fileInfo.Sys().(*syscall.Stat_t)
		if !ok {
			return fmt.Errorf("failed to get raw syscall.Stat_t data for '%s'", sourcePath)
		}

		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			dstPath, err := ensureOutputPath(dstPath)
			if err != nil {
				return err
			}
			if err := copyDir(sourcePath, dstPath); err != nil {
				return err
			}
		case os.ModeSymlink:
			panic("symlinks not supported")
		default:
			if err := copyFile(sourcePath, dstPath); err != nil {
				return err
			}
		}

		if err := os.Lchown(dstPath, int(stat.Uid), int(stat.Gid)); err != nil {
			return err
		}

		fInfo, err := entry.Info()
		if err != nil {
			return err
		}

		isSymlink := fInfo.Mode()&os.ModeSymlink != 0
		if !isSymlink {
			if err := os.Chmod(dstPath, fInfo.Mode()); err != nil {
				return err
			}
		}
	}
	return nil
}

func copyStaticFiles(manifest *Manifest, outputPath string) error {
	outputPath = fmt.Sprintf("%s/static/", outputPath)
	outputPath, err := ensureOutputPath(outputPath)
	if err != nil {

		return err
	}
	for _, path := range manifest.ManifestData.Static {
		fileInfo, err := os.Stat(path)
		if err != nil {
			return err
		}

		if fileInfo.IsDir() {
			err := copyDir(path, outputPath)
			if err != nil {
				return err
			}
		} else {
			err := copyFile(path, outputPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
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

func gen(path, recipesPath, manifestPath, outputPath string) error {
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

	manifest, err := loadManifest(manifestPath)
	if err != nil {
		return err
	}

	fmt.Printf("Copying static files: %v\n", manifest)
	return copyStaticFiles(manifest, outputPath)
}
