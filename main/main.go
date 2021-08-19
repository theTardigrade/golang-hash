package main

import (
	"fmt"

	hash "github.com/theTardigrade/golang-hash"
)

func main() {
	input := "a"
	fmt.Println(hash.Uint64String(input))
	fmt.Println(hash.Uint32String(input))
	fmt.Println()
	fmt.Println(hash.Uint16String(input))
	fmt.Println(uint16(hash.Uint32String(input)))
	fmt.Println(uint16(hash.Uint32String(input) >> 16))
	fmt.Println()
	fmt.Println(hash.Uint8String(input))
	fmt.Println(uint8(hash.Uint16String(input)))
	fmt.Println(uint8(hash.Uint16String(input) >> 8))
	fmt.Println()

	input = "stratum"
	fmt.Println(hash.Uint64String(input))
	fmt.Println(hash.Uint32String(input))
	fmt.Println()
	fmt.Println(hash.Uint16String(input))
	fmt.Println(uint16(hash.Uint32String(input)))
	fmt.Println(uint16(hash.Uint32String(input) >> 16))
	fmt.Println()
	fmt.Println(hash.Uint8String(input))
	fmt.Println(uint8(hash.Uint16String(input)))
	fmt.Println(uint8(hash.Uint16String(input) >> 8))
	fmt.Println()

	input = "test"
	fmt.Println(hash.Uint64String(input))
	fmt.Println(hash.Uint32String(input))
	fmt.Println()
	fmt.Println(hash.Uint16String(input))
	fmt.Println(uint16(hash.Uint32String(input)))
	fmt.Println(uint16(hash.Uint32String(input) >> 16))
	fmt.Println()
	fmt.Println(hash.Uint8String(input))
	fmt.Println(uint8(hash.Uint16String(input)))
	fmt.Println(uint8(hash.Uint16String(input) >> 8))
}
