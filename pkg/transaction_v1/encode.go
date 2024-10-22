package transactionv1

import (
	"encoding/binary"
	"errors"
	"io"
	"money_app/pkg/appconfig"
	"time"

	"github.com/shopspring/decimal"
)

var errCorruptedData = errors.New("corrupted transaction")

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

func (t Transaction) toBytes(config appconfig.Config) []byte {
	var data []byte

	// Store amount
	amountbytes := make([]byte, 8)
	binary.BigEndian.PutUint64(amountbytes, uint64(t.Amount.Shift(int32(config.Shift)).IntPart()))
	data = append(data, amountbytes...)

	// Store date
	data = append(data, dateToBytes(t.Date)...)

	// Store comment length and comment
	commentBytes := []byte(t.Comment)
	lengthBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes, uint32(len(commentBytes)))
	data = append(append(data, lengthBytes...), commentBytes...)

	return data
}

func (t Transaction) WriteToWriter(w io.Writer, config appconfig.Config) (int, error) {
	return w.Write(t.toBytes(config))
}

func (t *Transaction) ReadFromReader(r io.Reader, config appconfig.Config) (int, error) {
	data := make([]byte, 16)
	n, err := r.Read(data)
	if err != nil {
		if err == io.EOF {
			if n == 0 {
				return 0, err
			}
			return n, errCorruptedData
		}
		return n, err
	}

	t.Amount = decimal.New(int64(binary.BigEndian.Uint64(data[:8])), -int32(config.Shift))
	t.Date = dateFromBytes(data[8:12])

	commentLength := int(binary.BigEndian.Uint32(data[12:16]))
	data = make([]byte, commentLength)
	n, err = r.Read(data)
	if err != nil {
		if err == io.EOF {
			if n != int(commentLength) {
				return 16 + n, errCorruptedData
			}
		} else {
			return 16 + n, err
		}
	}
	t.Comment = string(data)
	return 16 + n, nil
}
