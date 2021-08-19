package hash

func Int32(data []byte) (hash int32) {
	hash = int32(Uint32(data))

	return
}

func Int32String(data string) (hash int32) {
	hash = Int32([]byte(data))

	return
}
