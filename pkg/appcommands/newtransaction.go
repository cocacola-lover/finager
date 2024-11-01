package appcommands

import (
	"fmt"
	"money_app/pkg/appconfig"
	"money_app/pkg/maputils"
	transactionv1 "money_app/pkg/transaction_v1"
	"os"
	"time"

	"github.com/peterh/liner"
	"github.com/shopspring/decimal"
)

func NewTransactionCommand(line *liner.State, config appconfig.Config, ctx time.Time) {
	strAmount, err := line.Prompt("Amount: ")
	if err != nil {
		fmt.Println("Failed to read amount: ", err)
		return
	}
	num, err := decimal.NewFromString(strAmount)
	if err != nil {
		fmt.Println("Failed to read amount: ", err)
		return
	}

	tag, err := line.Prompt("Tag: ")
	if err != nil {
		fmt.Println("Failed to read tag: ", err)
		return
	}
	tagHash, ok := maputils.FindKey(config.Tags, tag)
	if !ok {
		fmt.Println("Failed to read tag: unknown tag")
		return
	}

	comment, err := line.Prompt("Comment: ")
	if err != nil {
		fmt.Println("Failed to read comment: ", err)
		return
	}

	newT := transactionv1.Transaction{
		Amount:  num,
		Date:    ctx,
		Comment: comment,
		Tag:     tagHash,
	}

	if err := os.MkdirAll("transaction-history", 0755); err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	file, err := os.OpenFile(fmt.Sprintf("transaction-history/%d.%d.bin", ctx.Month(), ctx.Year()), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = newT.WriteToWriter(file, config)
	if err != nil {
		fmt.Println("Error writing transaction:", err)
		return
	}
	fmt.Println("Added transaction")
}
