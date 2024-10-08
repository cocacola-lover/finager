package appcommands

import (
	"fmt"
	"io"
	"money_app/pkg/appconfig"
	transactionv1 "money_app/pkg/transaction_v1"
	"os"

	"github.com/peterh/liner"
)

func ReadTransactionCommand(line *liner.State, config appconfig.Config) error {
	file, err := os.Open("transaction-history")
	if err != nil {
		return err
	}
	defer file.Close()

	for {
		var t transactionv1.Transaction

		bytes := make([]byte, 12)
		_, err := file.Read(bytes)

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}
		if err := t.FromBytes(bytes, config); err != nil {
			return err
		}

		fmt.Println("Decoded:", t)
	}
	return nil
}
