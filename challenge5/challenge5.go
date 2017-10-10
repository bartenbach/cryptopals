package challenge5

// RepeatingXOR executes a repeating XOR of the supplied array of
// bytes and the supplied key, returning the result
func RepeatingXOR(value, key []byte) []byte {
	result := make([]byte, len(value))
	for i := range value {
		result[i] = value[i] ^ key[i%len(key)]
	}
	return result
}
