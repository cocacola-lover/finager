package transactionv1

import (
	"encoding/binary"
	"io"
	"money_app/pkg/appconfig"
	"money_app/pkg/apperrors"
	"money_app/pkg/encodingutils"
	"time"

	"github.com/shopspring/decimal"
)

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

func (t Transaction) toBytes(config appconfig.Config) ([]byte, error) {
	var data []byte

	// Store amount
	data = append(data, encodingutils.Uint64ToBytes(uint64(t.Amount.Shift(int32(config.Shift)).IntPart()))...)

	// Store date
	data = append(data, dateToBytes(t.Date)...)

	// Store tag
	data = append(data, encodingutils.Uint32ToBytes(t.Tag)...)

	// Store comment length and comment
	data = append(data, encodingutils.StringToBytes(t.Comment)...)

	return data, nil
}

func (t Transaction) WriteToWriter(w io.Writer, config appconfig.Config) (int, error) {
	bytes, err := t.toBytes(config)
	if err != nil {
		return 0, err
	}
	return w.Write(bytes)
}

func (t *Transaction) ReadFromReader(r io.Reader, config appconfig.Config) (int, error) {
	data := make([]byte, 20)
	n, err := r.Read(data)
	if err != nil {
		if err == io.EOF {
			if n == 0 {
				return 0, err
			}
			return n, apperrors.ErrCorruptedData
		}
		return n, err
	}

	t.Amount = decimal.New(
		int64(binary.BigEndian.Uint64(data[:8])),
		-int32(config.Shift),
	)
	t.Date = dateFromBytes(data[8:12])

	t.Tag = binary.BigEndian.Uint32(data[12:16])
	if _, ok := config.Tags[t.Tag]; !ok {
		return 20, apperrors.ErrCorruptedData
	}

	commentLength := int(binary.BigEndian.Uint32(data[16:20]))
	data = make([]byte, commentLength)
	n, err = r.Read(data)
	if err != nil {
		if err == io.EOF {
			if n != int(commentLength) {
				return 20 + n, apperrors.ErrCorruptedData
			}
		} else {
			return 20 + n, err
		}
	}
	t.Comment = string(data)
	return 20 + n, nil
}
