package config

import (
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"time"
)

type Config struct {
	UrlTemplate string        `yaml:"url_template"`
	Teams       []string      `yaml:"teams"`
	Timeout     time.Duration `yaml:"timeout"`
}

func Parse(reader io.Reader) (*Config, error) {
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var c Config
	if err := yaml.Unmarshal(b, &c); err != nil {
		return nil, err
	}
	return &c, nil
}
