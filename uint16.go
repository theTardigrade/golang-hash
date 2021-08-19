package hash

func Uint16(data []byte) (hash uint16) {
	hash32 := Uint32(data)

	hash = uint16(hash32) ^ uint16(hash32>>16)

	return
}

func Uint16String(data string) (hash uint16) {
	hash = Uint16([]byte(data))

	return
}
