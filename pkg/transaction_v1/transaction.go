package transactionv1

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	Amount decimal.Decimal
	Date   time.Time
}
