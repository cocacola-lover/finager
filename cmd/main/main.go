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

cmdLoop:
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

		switch cmd {
		case "NEW":
			appcommands.NewTransactionCommand(line, config, timeContext)
		case "READ":
			appcommands.ReadTransactionCommand(line, config)
		case "GETDATE":
			appcontext.DisplayTime(timeContext)
		case "SETDATE":
			if err = appcontext.ParseTime(line, &timeContext); err != nil {
				fmt.Println("Failed to set date: ", err.Error())
			}
		case "ADDTAG":
			appcommands.AddTag(line, &config)
		case "READTAGS":
			appcommands.ReadTags(config)
		case "QUIT":
			break cmdLoop
		default:
			fmt.Println("Unknown command:", cmd)
		}
	}
}
