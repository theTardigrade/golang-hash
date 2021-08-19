package hash

func Int64(data []byte) (hash int64) {
	hash = int64(Uint64(data))

	return
}

func Int64String(data string) (hash int64) {
	hash = Int64([]byte(data))

	return
}
