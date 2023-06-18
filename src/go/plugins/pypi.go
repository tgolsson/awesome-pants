package plugins

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type PackageInfo struct {
	Author                 string      `json:"author"`
	AuthorEmail            string      `json:"author_email"`
	BugtrackURL            interface{} `json:"bugtrack_url"`
	Classifiers            []string    `json:"classifiers"`
	Description            string      `json:"description"`
	DescriptionContentType string      `json:"description_content_type"`
	DocsURL                interface{} `json:"docs_url"`
	DownloadURL            string      `json:"download_url"`
	Downloads              struct {
		LastDay   int `json:"last_day"`
		LastMonth int `json:"last_month"`
		LastWeek  int `json:"last_week"`
	} `json:"downloads"`
	HomePage        string            `json:"home_page"`
	Keywords        string            `json:"keywords"`
	License         string            `json:"license"`
	Maintainer      string            `json:"maintainer"`
	MaintainerEmail string            `json:"maintainer_email"`
	Name            string            `json:"name"`
	PackageURL      string            `json:"package_url"`
	Platform        interface{}       `json:"platform"`
	ProjectURL      string            `json:"project_url"`
	ProjectUrls     map[string]string `json:"project_urls"`
	ReleaseURL      string            `json:"release_url"`
	RequiresDist    []string          `json:"requires_dist"`
	RequiresPython  string            `json:"requires_python"`
	Summary         string            `json:"summary"`
	Version         string            `json:"version"`
	Yanked          bool              `json:"yanked"`
	YankedReason    interface{}       `json:"yanked_reason"`
}

type Digests struct {
	Blake2B256 string `json:"blake2b_256"`
	Md5        string `json:"md5"`
	Sha256     string `json:"sha256"`
}

type ReleaseInfo struct {
	CommentText       string      `json:"comment_text"`
	Digests           Digests     `json:"digests"`
	Downloads         int         `json:"downloads"`
	Filename          string      `json:"filename"`
	HasSig            bool        `json:"has_sig"`
	Md5Digest         string      `json:"md5_digest"`
	PackageType       string      `json:"packagetype"`
	PythonVersion     string      `json:"python_version"`
	RequiresPython    interface{} `json:"requires_python"`
	Size              int         `json:"size"`
	UploadTime        string      `json:"upload_time"`
	UploadTimeIso8601 time.Time   `json:"upload_time_iso_8601"`
	URL               string      `json:"url"`
	Yanked            bool        `json:"yanked"`
	YankedReason      interface{} `json:"yanked_reason"`
}

type UrlInfo struct {
	CommentText       string      `json:"comment_text"`
	Digests           Digests     `json:"digests"`
	Downloads         int         `json:"downloads"`
	Filename          string      `json:"filename"`
	HasSig            bool        `json:"has_sig"`
	Md5Digest         string      `json:"md5_digest"`
	PackageType       string      `json:"packagetype"`
	PythonVersion     string      `json:"python_version"`
	RequiresPython    any         `json:"requires_python"`
	Size              int         `json:"size"`
	UploadTime        string      `json:"upload_time"`
	UploadTimeIso8601 time.Time   `json:"upload_time_iso_8601"`
	URL               string      `json:"url"`
	Yanked            bool        `json:"yanked"`
	YankedReason      interface{} `json:"yanked_reason"`
}

type Package struct {
	Info       PackageInfo              `json:"info"`
	Releases   map[string][]ReleaseInfo `json:"releases"`
	LastSerial int                      `json:"last_serial"`
	Urls       []UrlInfo                `json:"urls"`
}

func getPackageInfo(packageName string) (*Package, error) {
	url := fmt.Sprintf("https://pypi.org/pypi/%s/json", packageName)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var packageInfo Package
	err = json.NewDecoder(resp.Body).Decode(&packageInfo)
	if err != nil {
		return nil, err
	}

	return &packageInfo, nil
}
