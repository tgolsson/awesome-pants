package main

import (
	"fmt"

	"github.com/tgolsson/awesome-pants/src/go/plugins"
)

func refresh(path, output string) error {
	data, err := plugins.LoadCollection(path)
	if err != nil {
		return err
	}

	for i, plugin := range data.Plugins {
		fmt.Printf("Refreshing %s\n", plugin.Name)
		err := plugins.RefreshPlugin(&plugin)
		if err != nil {
			return err
		}

		fmt.Printf("Done refreshing %s\n", plugin.Name)
		data.Plugins[i] = plugin
	}

	return plugins.WriteCollection(output, data)
}
