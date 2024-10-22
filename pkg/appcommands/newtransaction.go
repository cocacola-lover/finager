package appcommands

import (
	"fmt"
	"money_app/pkg/appconfig"
	transactionv1 "money_app/pkg/transaction_v1"
	"os"
	"time"

	"github.com/peterh/liner"
	"github.com/shopspring/decimal"
)

func NewTransactionCommand(line *liner.State, config appconfig.Config, ctx time.Time) error {
	strAmount, err := line.Prompt("Amount: ")
	if err != nil {
		return err
	}
	num, err := decimal.NewFromString(strAmount)
	if err != nil {
		return err
	}

	comment, err := line.Prompt("Comment: ")
	if err != nil {
		return err
	}

	newT := transactionv1.Transaction{
		Amount:  num,
		Date:    ctx,
		Comment: comment,
	}

	file, err := os.OpenFile("transaction-history", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	_, err = newT.WriteToWriter(file, config)
	return err
}
