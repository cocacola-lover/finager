package appconfig

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ReadConfig() (Config, error) {
	var config Config

	data, err := os.ReadFile("config.yaml")
	yaml.Unmarshal(data, &config)

	return config, err
}
