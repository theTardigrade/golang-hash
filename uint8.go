package hash

func Uint8(data []byte) (hash uint8) {
	hash16 := Uint16(data)

	hash = uint8(hash16) ^ uint8(hash16>>8)

	return
}

func Uint8String(data string) (hash uint8) {
	hash = Uint8([]byte(data))

	return
}
