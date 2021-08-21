package hash

import (
	"math/big"
)

const (
	int128MinString = "-0x80000000000000000000000000000000"
	int128ModString = "0x80000000000000000000000000000000"
)

var (
	int128Min = new(big.Int)
	int128Mod = new(big.Int)
)

func init() {
	int128Min.SetString(int128MinString, 0)
	int128Mod.SetString(int128ModString, 0)
}

func Int128(data []byte) (hash *big.Int) {
	hash = Uint128(data)

	delta := new(big.Int)
	delta.Sub(hash, int128Mod)

	zero := big.NewInt(0)

	if delta.Cmp(zero) >= 0 {
		hash.Add(int128Min, delta)
	}

	return
}

func Int128String(data string) (hash *big.Int) {
	hash = Int128([]byte(data))

	return
}
