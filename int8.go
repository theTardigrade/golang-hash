package hash

func Int8(data []byte) (hash int8) {
	hash = int8(Uint8(data))

	return
}

func Int8String(data string) (hash int8) {
	hash = Int8([]byte(data))

	return
}
