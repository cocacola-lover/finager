package encodingutils

import "encoding/binary"

func Uint64ToBytes(val uint64) (ans []byte) {
	ans = make([]byte, 8)
	binary.BigEndian.PutUint64(ans, val)
	return
}

func Uint32ToBytes(val uint32) (ans []byte) {
	ans = make([]byte, 4)
	binary.BigEndian.PutUint32(ans, val)
	return
}
