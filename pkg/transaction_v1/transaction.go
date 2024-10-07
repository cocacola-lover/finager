package transactionv1

import "time"

type Transaction struct {
	Amount float32
	Date   time.Time
}
