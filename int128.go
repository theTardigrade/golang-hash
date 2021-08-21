package hash

import (
	"math/big"
)

const (
	int128MinString = "-0x80000000000000000000000000000000"
	int128ModString = "0x80000000000000000000000000000000"
)

func Int128(data []byte) (hash *big.Int) {
	hash = Uint128(data)

	mod := new(big.Int)
	mod.SetString(int128ModString, 0)

	delta := new(big.Int)
	delta.Sub(hash, mod)

	zero := big.NewInt(0)

	if delta.Cmp(zero) >= 0 {
		min := new(big.Int)
		min.SetString(int128MinString, 0)

		hash.Add(min, delta)
	}

	return
}

func Int128String(data string) (hash *big.Int) {
	hash = Int128([]byte(data))

	return
}
