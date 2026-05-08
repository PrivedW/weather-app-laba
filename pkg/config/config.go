package config

import (
	"bytes"
	_ "embed"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

//go:embed default_config.yaml
var defaultConfig []byte

type ConfigFile struct {
	C Config `yaml:"service"`
}
type Provider struct {
	Type string `yaml:"type"`
}
type Location struct {
	Lat  float64 `yaml:"lat"`
	Long float64 `yaml:"long"`
}
type Config struct {
	P Provider `yaml:"provider"`
	L Location `yaml:"location"`
}

func Parse(r io.Reader) (Config, error) {
	var c ConfigFile
	if err := yaml.NewDecoder(r).Decode(&c); err != nil {
		return Config{}, err
	}
	return c.C, nil
}

func ParseDefault() (Config, error) {
	return Parse(bytes.NewReader(defaultConfig))
}

func ParsePath(path string) (Config, error) {
	r, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer func() {
		_ = r.Close()
	}()

	return Parse(r)
}
