package hash

import (
	"math/big"
)

var (
	int256Min = new(big.Int)
	int256Mod = new(big.Int)
)

func init() {
	int256Mod.Lsh(bigOne, 255)
	int256Min.Neg(int256Mod)
}

func Int256(data []byte) (hash *big.Int) {
	hash = Uint256(data)

	delta := new(big.Int)
	delta.Sub(hash, int256Mod)

	if delta.Cmp(bigZero) >= 0 {
		hash.Add(int256Min, delta)
	}

	return
}

func Int256String(data string) (hash *big.Int) {
	hash = Int256([]byte(data))

	return
}
