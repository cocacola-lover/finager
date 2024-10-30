package main

import (
	"fmt"
	"money_app/pkg/appcommands"
	"money_app/pkg/appconfig"
	"money_app/pkg/appcontext"
	"strings"
	"time"

	"github.com/peterh/liner"
)

func promptString(timeContext time.Time) string {
	return fmt.Sprintf("|%s %d|> ", timeContext.Month().String(), timeContext.Day())
}

func main() {

	config, err := appconfig.ReadConfig()
	if err != nil {
		fmt.Println("Failed to read config. Exiting...")
		return
	}

	timeContext := time.Now()

	line := liner.NewLiner()
	defer line.Close()

	commands := []string{"NEW", "READ"}

	line.SetCompleter(func(line string) (c []string) {
		for _, n := range commands {
			if strings.HasPrefix(strings.ToUpper(n), strings.ToUpper(line)) {
				c = append(c, n)
			}
		}
		return
	})

	line.SetMultiLineMode(true)
	line.SetCtrlCAborts(true)
	line.SetTabCompletionStyle(liner.TabPrints)

	for {
		cmd, err := line.Prompt(promptString(timeContext))
		if err != nil {
			if err == liner.ErrPromptAborted {
				fmt.Println("Aborted")
				continue
			}
			fmt.Println("Error reading input:", err)
			return
		}

		line.AppendHistory(cmd)
		if cmd == "NEW" {
			appcommands.NewTransactionCommand(line, config, timeContext)
		} else if cmd == "READ" {
			err = appcommands.ReadTransactionCommand(line, config)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else if cmd == "GETDATE" {
			appcontext.DisplayTime(timeContext)
		} else if cmd == "SETDATE" {
			if err = appcontext.ParseTime(line, &timeContext); err != nil {
				fmt.Println("Failed to set date: ", err.Error())
			}
		} else if cmd == "ADDTAG" {
			appcommands.AddTag(line, &config)
		} else if cmd == "READTAGS" {
			appcommands.ReadTags(config)
		} else if cmd == "quit" {
			break
		} else {
			fmt.Println("Unknown command:", cmd)
		}
	}
}
