package appcommands

import (
	"fmt"
	"money_app/pkg/appconfig"
)

func ReadTags(config appconfig.Config) {
	for _, v := range config.Tags {
		fmt.Println(v)
	}
}
