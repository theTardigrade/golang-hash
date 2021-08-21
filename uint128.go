package hash

import (
	"math/big"
)

const (
	uint128ModString    = "0x100000000000000000000000000000000"
	uint128PrimeString  = "0x0000000001000000000000000000013b"
	uint128OffsetString = "0x6c62272e07bb014262b821756295c58d"
)

func Uint128(data []byte) (hash *big.Int) {
	hash = new(big.Int)
	hash.SetString(uint128OffsetString, 0)

	prime := new(big.Int)
	prime.SetString(uint128PrimeString, 0)

	mod := new(big.Int)
	mod.SetString(uint128ModString, 0)

	datum := new(big.Int)
	for _, d := range data {
		hash.Xor(hash, datum.SetBytes([]byte{d}))
		hash.Mul(hash, prime)
		hash.Mod(hash, mod)
	}

	return
}

func Uint128String(data string) (hash *big.Int) {
	hash = Uint128([]byte(data))

	return
}
