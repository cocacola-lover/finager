package appcommands

import (
	"fmt"
	"io"
	"money_app/pkg/appconfig"
	transactionv1 "money_app/pkg/transaction_v1"
	"os"
	"sort"
	"time"

	"github.com/peterh/liner"
)

func printDate(date time.Time) {
	fmt.Printf("-------\n\033[31m* %d/%d/%d\033[0m\n-------\n", date.Day(), date.Month(), date.Year())
}

func readFileIntoArr(fileName string, transactionArr *[]transactionv1.Transaction, config appconfig.Config) error {

	file, err := os.Open(fmt.Sprintf("transaction-history/%s", fileName))
	if err != nil {
		return err
	}
	defer file.Close()

	for {
		var t transactionv1.Transaction
		_, err = t.ReadFromReader(file, config)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		*transactionArr = append(*transactionArr, t)
	}

	return nil
}

func ReadTransactionCommand(line *liner.State, config appconfig.Config) {

	var transactionArr []transactionv1.Transaction

	entries, err := os.ReadDir("transaction-history")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, entry := range entries {
		err = readFileIntoArr(entry.Name(), &transactionArr, config)
		if err != nil {
			fmt.Printf("Error reading file %s : %s\n", entry.Name(), err.Error())
			return
		}
	}

	if len(transactionArr) == 0 {
		return
	}

	sort.Slice(transactionArr, func(i, j int) bool {
		return transactionArr[j].Date.Before(transactionArr[i].Date)
	})

	curDate := transactionArr[0].Date
	printDate(curDate)
	for _, t := range transactionArr {
		if !t.Date.Equal(curDate) {
			curDate = t.Date
			printDate(curDate)
		}
		fmt.Printf("Spent %v₽ | %s | %s\n", t.Amount, config.Tags[t.Tag], t.Comment)
	}
}
