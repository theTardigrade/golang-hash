package hash

import (
	"math/big"
)

const (
	uint128ModString    = "0x100000000000000000000000000000000"
	uint128PrimeString  = "0x0000000001000000000000000000013b"
	uint128OffsetString = "0x6c62272e07bb014262b821756295c58d"
)

var (
	uint128Mod    = new(big.Int)
	uint128Prime  = new(big.Int)
	uint128Offset = new(big.Int)
)

func init() {
	uint128Mod.SetString(uint128ModString, 0)
	uint128Prime.SetString(uint128PrimeString, 0)
	uint128Offset.SetString(uint128OffsetString, 0)
}

func Uint128(data []byte) (hash *big.Int) {
	hash = new(big.Int)
	hash.Set(uint128Offset)

	datum := new(big.Int)
	for _, d := range data {
		hash.Xor(hash, datum.SetBytes([]byte{d}))
		hash.Mul(hash, uint128Prime)
		hash.Mod(hash, uint128Mod)
	}

	return
}

func Uint128String(data string) (hash *big.Int) {
	hash = Uint128([]byte(data))

	return
}
