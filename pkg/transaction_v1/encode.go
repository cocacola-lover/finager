package transactionv1

import (
	"encoding/binary"
	"math"
	"time"
)

func float32ToBytes(num float32) []byte {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, math.Float32bits(num))
	return bytes
}

// Requires 4 bytes
func bytesToFloat32(bytes []byte) float32 {
	bits := binary.BigEndian.Uint32(bytes)
	f := math.Float32frombits(bits)
	return f
}
func uint16ToBytes(num uint16) []byte {
	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, num)
	return bytes
}
func bytesToUint16(bytes []byte) uint16 {
	return binary.BigEndian.Uint16(bytes)
}
func dateToBytes(date time.Time) []byte {
	year, month, day := date.Date()
	return append(
		uint16ToBytes(uint16(year)),
		uint8(month),
		uint8(day),
	)
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
	return append(
		float32ToBytes(t.Amount),
		dateToBytes(t.Date)...,
	), nil
}

// Requires 8 bytes
func (t *Transaction) FromBytes(data []byte) error {
	t.Amount = bytesToFloat32(data[0:4])
	t.Date = dateFromBytes(data[4:])
	return nil
}
