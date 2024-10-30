package maputils

func FindKey[K comparable, V comparable](m map[K]V, val V) (ans K, ok bool) {
	for k, v := range m {
		if v == val {
			ans, ok = k, true
			return
		}
	}
	return
}
