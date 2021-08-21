package hash

import (
	"math/big"
)

var (
	int128Min = new(big.Int)
	int128Mod = new(big.Int)
)

func init() {
	int128Mod.Lsh(bigOne, 127)
	int128Min.Neg(int128Mod)
}

func Int128(data []byte) (hash *big.Int) {
	hash = Uint128(data)

	delta := new(big.Int)
	delta.Sub(hash, int128Mod)

	if delta.Cmp(bigZero) >= 0 {
		hash.Add(int128Min, delta)
	}

	return
}

func Int128String(data string) (hash *big.Int) {
	hash = Int128([]byte(data))

	return
}
