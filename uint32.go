package hash

const (
	uint32Prime  uint32 = 0x01000193
	uint32Offset uint32 = 0x811c9dc5
)

func Uint32(data []byte) (hash uint32) {
	hash = uint32Offset

	for _, b := range data {
		hash ^= uint32(b)
		hash *= uint32Prime
	}

	return
}

func Uint32String(data string) (hash uint32) {
	hash = Uint32([]byte(data))

	return
}
