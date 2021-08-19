package hash

const (
	uint64Prime  uint64 = 0x00000100000001b3
	uint64Offset uint64 = 0xcbf29ce484222325
)

func Uint64(data []byte) (hash uint64) {
	hash = uint64Offset

	for _, b := range data {
		hash ^= uint64(b)
		hash *= uint64Prime
	}

	return
}

func Uint64String(data string) (hash uint64) {
	hash = Uint64([]byte(data))

	return
}
