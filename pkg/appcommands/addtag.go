package appcommands

import (
	"fmt"
	"money_app/pkg/appconfig"
	"money_app/pkg/apptags"

	"github.com/peterh/liner"
)

func AddTag(line *liner.State, config *appconfig.Config) {
	tag, err := line.Prompt("Tag: ")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = apptags.AddTag(tag)
	if err != nil {
		fmt.Println(err.Error())
	}

	*config, err = appconfig.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Added tag")
	}
}
