package feed

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
	"github.com/mmcdole/gofeed"
)

type Feed struct {
	Url      string
	Articles []Article
}

type Article struct {
	Title string
	Url   string
	Date  string
}

func ParseYaml(yamlFilePath string) ([]Feed, error) {
	absolutePath, err := createAbsolutePath(yamlFilePath)
	if err != nil {
		return nil, err
	}

	buf, err := os.ReadFile(absolutePath)
	if err != nil {
		return nil, err
	}

	var urls []string
	err = yaml.Unmarshal(buf, &urls)
	if err != nil {
		return nil, err
	}

	var feeds []Feed
	for _, url := range urls {
		feeds = append(feeds, Feed{Url: url})
	}

	return feeds, nil
}

func ParseRss(feeds []Feed) {
	for _, feed := range feeds {
		item, _ := gofeed.NewParser().ParseURL(feed.Url)
		fmt.Println(item)
	}
}

func createAbsolutePath(path string) (string, error) {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	return absolutePath, nil
}
