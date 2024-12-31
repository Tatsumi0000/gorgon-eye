package config

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	YamlFileName string
	Urls         []string `yaml:"urls"`
}

func New(yamlFileName string) *Config {
	return &Config{
		YamlFileName: yamlFileName,
	}
}

func (c *Config) Read() {
	buf, err := os.ReadFile(c.YamlFileName)

	if err != nil {
		fmt.Println(err)
		return
	}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(c.Urls)
}
