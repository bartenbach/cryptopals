package challenge2

// XOR performs an XOR of two byte arrays together, returning the resulting byte array
func XOR(val1, val2 []byte) []byte {
	if len(val1) != len(val2) {
		panic("Lengths of byte arrays not equal!")
	}
	res := make([]byte, len(val1))
	for i := range val1 {
		res[i] = val1[i] ^ val2[i]
	}
	return res
}
