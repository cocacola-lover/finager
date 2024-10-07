package main

import (
	"encoding/gob"
	"fmt"
	"money_app/pkg/appcommands"
	transactionv1 "money_app/pkg/transaction_v1"
	"strings"

	"github.com/peterh/liner"
)

func main() {
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

	gob.Register(transactionv1.Transaction{})

	for {
		cmd, err := line.Prompt("gosql> ")
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
			err = appcommands.NewTransactionCommand(line)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Added transaction")
			}
		} else if cmd == "READ" {
			err = appcommands.ReadTransactionCommand(line)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Shown transaction")
			}
		} else if cmd == "quit" {
			break
		} else {
			fmt.Println("Unknown command:", cmd)
		}
	}

	fmt.Println("Closed")
}
