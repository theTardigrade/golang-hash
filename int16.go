package hash

func Int16(data []byte) (hash int16) {
	hash = int16(Uint16(data))

	return
}

func Int16String(data string) (hash int16) {
	hash = Int16([]byte(data))

	return
}
