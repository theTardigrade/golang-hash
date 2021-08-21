package hash

import (
	"math/big"
)

var (
	int256Min = new(big.Int)
	int256Mod = new(big.Int)
)

func init() {
	int256Mod.Lsh(big.NewInt(1), 255)
	int256Min.Neg(int256Mod)
}

func Int256(data []byte) (hash *big.Int) {
	hash = Uint256(data)

	delta := new(big.Int)
	delta.Sub(hash, int256Mod)

	zero := big.NewInt(0)

	if delta.Cmp(zero) >= 0 {
		hash.Add(int256Min, delta)
	}

	return
}

func Int256String(data string) (hash *big.Int) {
	hash = Int256([]byte(data))

	return
}
