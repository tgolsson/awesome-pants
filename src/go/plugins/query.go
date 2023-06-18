package plugins

// Simple package to query plugin metadata from PyPi and Github.

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/BurntSushi/toml"
)

type Metadata struct {
	Author      string `toml:"author"`
	AuthorEmail string `toml:"author_email"`
	Website     string `toml:"website"`
	Summary     string `toml:"summary"`

	LatestVersion string   `toml:"latest_version"`
	Tags          []string `toml:"tags"`

	Versions []Version `toml:"versions"`
	License  string    `toml:"license"`
}

type DownloadStatistics struct {
	Yesterday int `toml:"last_day" json:"last_day"`
	Week      int `toml:"last_week" json:"last_weeK"`
	Month     int `toml:"last_month" json:"last_month"`
}

type Plugin struct {
	Name     string             `toml:"name"`
	Pypi     string             `toml:"pypi"`
	Metadata Metadata           `toml:"metadata"`
	Stats    DownloadStatistics `toml:"stats"`
	Status   string             `toml:"status"`
	Icon     string             `toml:"icon"`
}

type Version struct {
	Version     string `toml:"version"`
	ReleaseDate string `toml:"release_date"`
}

type downloadStatisticsResponse struct {
	Data    DownloadStatistics `toml:"data"`
	Type    string             `toml:"type"`
	Package string             `toml:"package"`
}

type PluginsCache struct {
	Plugins []Plugin `toml:"plugins"`
}

func retrieveDownloadStatistics(plugin *Plugin) (*DownloadStatistics, error) {
	// Retrieve download statistics from PyPi
	url := fmt.Sprintf("https://pypistats.org/api/packages/%s/recent", plugin.Name)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error retrieving download statistics for plugin %s: %s", plugin.Name, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body for plugin %s: %s", plugin.Name, err)
	}
	var stats downloadStatisticsResponse
	err = json.Unmarshal(body, &stats)
	if err != nil {

		return nil, fmt.Errorf("error unmarshalling response body for plugin %s: %s", plugin.Name, err)
	}

	return &stats.Data, nil
}

func RefreshPlugin(plugin *Plugin) error {
	packageInfo, err := getPackageInfo(plugin.Name)
	if err != nil {
		return fmt.Errorf("error retrieving package info for plugin %s: %s", plugin.Name, err)
	}

	downloadStatistics, err := retrieveDownloadStatistics(plugin)
	if err != nil {
		return fmt.Errorf("error retrieving download statistics for plugin %s: %s", plugin.Name, err)
	}

	plugin.Stats = *downloadStatistics

	if packageInfo.Info.Author != "" {
		plugin.Metadata.Author = packageInfo.Info.Author
		plugin.Metadata.AuthorEmail = packageInfo.Info.AuthorEmail
	} else {
		// If the author is not set, assume the email is on the form "name <email>" and parse it.
		author := strings.Split(packageInfo.Info.AuthorEmail, "<")
		if len(author) == 2 {
			plugin.Metadata.Author = strings.TrimSpace(author[0])
			plugin.Metadata.AuthorEmail = strings.TrimRight(strings.TrimSpace(author[1]), ">")
		} else {
			plugin.Metadata.Author = packageInfo.Info.Author
			plugin.Metadata.AuthorEmail = packageInfo.Info.AuthorEmail
		}
	}

	if packageInfo.Info.HomePage != "" {
		plugin.Metadata.Website = packageInfo.Info.HomePage
	} else if packageInfo.Info.ProjectURL != "" {
		plugin.Metadata.Website = packageInfo.Info.ProjectURL
	} else {
		maybeUrls := []string{"Homepage", "Repository", "Website"}
		for _, url := range maybeUrls {
			if val, ok := packageInfo.Info.ProjectUrls[url]; ok {
				plugin.Metadata.Website = val
				break
			}
		}
	}

	plugin.Metadata.Summary = packageInfo.Info.Summary
	plugin.Metadata.LatestVersion = packageInfo.Info.Version
	plugin.Metadata.Tags = strings.Split(packageInfo.Info.Keywords, ",")
	plugin.Metadata.License = packageInfo.Info.License

	for version, files := range packageInfo.Releases {
		for _, data := range files {
			if data.PackageType == "bdist_wheel" {
				plugin.Metadata.Versions = append(plugin.Metadata.Versions, Version{
					ReleaseDate: data.UploadTime,
					Version:     version,
				})
				break
			}
		}
	}

	sort.Slice(plugin.Metadata.Versions, func(i, j int) bool {
		return plugin.Metadata.Versions[i].Version > plugin.Metadata.Versions[j].Version
	})

	return nil
}

// LoadCollection loads the plugin collection from the given path.
func LoadCollection(path string) (*PluginsCache, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data PluginsCache
	if _, err := toml.Decode(string(content), &data); err != nil {
		return nil, err
	}

	return &data, nil
}

// WriteCollection writes the plugin collection to the given path.
func WriteCollection(path string, data *PluginsCache) error {
	file, err := os.OpenFile(fmt.Sprintf("%s.tmp", path), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	err = toml.NewEncoder(file).Encode(data)
	if err != nil {
		file.Close()
		return err
	}

	file.Close()
	os.Rename(fmt.Sprintf("%s.tmp", path), path)

	return nil
}
