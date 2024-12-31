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

func (c *Config) Read() error {
	buf, err := os.ReadFile(c.YamlFileName)

	if err != nil {
		return err
	}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return err
	}
	fmt.Println(c.Urls)
	return nil
}
