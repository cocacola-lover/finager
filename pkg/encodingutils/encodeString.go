package encodingutils

func StringToBytes(val string) (ans []byte) {
	tempStore := []byte(val)
	ans = Uint32ToBytes(uint32(len(tempStore)))
	ans = append(ans, tempStore...)
	return
}
