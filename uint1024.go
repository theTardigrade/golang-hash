package hash

import (
	"math/big"
)

const (
	uint1024PrimeString  = "0x10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000018D"
	uint1024OffsetString = "0x5f7a76758ecc4d32e56d5a591028b74b29fc4223fdada16c3bf34eda3674da9a21d9000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004c6d7eb6e73802734510a555f256cc005ae556bde8cc9c6a93b21aff4b16c71ee90b3"
)

var (
	uint1024Mod    = new(big.Int)
	uint1024Prime  = new(big.Int)
	uint1024Offset = new(big.Int)
)

func init() {
	uint1024Mod.Lsh(bigOne, 1024)
	uint1024Prime.SetString(uint1024PrimeString, 0)
	uint1024Offset.SetString(uint1024OffsetString, 0)
}

func Uint1024(data []byte) (hash *big.Int) {
	hash = new(big.Int)
	hash.Set(uint1024Offset)

	datum := new(big.Int)
	for _, d := range data {
		hash.Xor(hash, datum.SetUint64(uint64(d)))
		hash.Mul(hash, uint1024Prime)
		hash.Mod(hash, uint1024Mod)
	}

	return
}

func Uint1024String(data string) (hash *big.Int) {
	hash = Uint1024([]byte(data))

	return
}
