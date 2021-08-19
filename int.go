package hash

import "strconv"

func Int(data []byte) (hash int) {
	switch strconv.IntSize {
	case 32:
		hash = int(Int32(data))
	default:
		hash = int(Int64(data))
	}

	return
}

func IntString(data string) (hash int) {
	hash = Int([]byte(data))

	return
}
