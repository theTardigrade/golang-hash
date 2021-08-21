package hash

import (
	"math/big"
)

var (
	int512Min = new(big.Int)
	int512Mod = new(big.Int)
)

func init() {
	int512Mod.Lsh(bigOne, 511)
	int512Min.Neg(int512Mod)
}

func Int512(data []byte) (hash *big.Int) {
	hash = Uint512(data)

	delta := new(big.Int)
	delta.Sub(hash, int512Mod)

	if delta.Cmp(bigZero) >= 0 {
		hash.Add(int512Min, delta)
	}

	return
}

func Int512String(data string) (hash *big.Int) {
	hash = Int512([]byte(data))

	return
}
