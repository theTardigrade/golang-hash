package hash

import "strconv"

func Uint(data []byte) (hash uint) {
	switch strconv.IntSize {
	case 32:
		hash = uint(Uint32(data))
	default:
		hash = uint(Uint64(data))
	}

	return
}

func UintString(data string) (hash uint) {
	hash = Uint([]byte(data))

	return
}
