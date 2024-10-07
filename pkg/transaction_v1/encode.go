package transactionv1

import (
	"encoding/binary"
	"time"

	"github.com/shopspring/decimal"
)

const SHIFT = 2

func dateToBytes(date time.Time) []byte {
	year, month, day := date.Date()

	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, uint16(year))

	return append(bytes, uint8(month), uint8(day))
}

// Requires 4bytes
func dateFromBytes(bytes []byte) time.Time {
	return time.Date(
		int(binary.BigEndian.Uint16(bytes[0:2])),
		time.Month(bytes[2]),
		int(bytes[3]),
		0, 0, 0, 0, time.Local,
	)
}

func (t Transaction) ToBytes() ([]byte, error) {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, uint64(t.Amount.Shift(SHIFT).IntPart()))

	return append(
		bytes,
		dateToBytes(t.Date)...,
	), nil
}

// Requires 12 bytes
func (t *Transaction) FromBytes(data []byte) error {
	t.Amount = decimal.New(int64(binary.BigEndian.Uint64(data[:8])), -SHIFT)
	t.Date = dateFromBytes(data[8:])
	return nil
}
