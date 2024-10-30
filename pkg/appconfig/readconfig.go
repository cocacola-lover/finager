package appconfig

import (
	"money_app/pkg/apptags"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadConfig() (Config, error) {
	var config Config

	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return Config{}, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	tags, err := apptags.ReadTags()
	if err != nil {
		return Config{}, err
	}

	config.Tags = tags
	return config, err
}
