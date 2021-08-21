package hash

import (
	"math/big"
)

const (
	int256MinString = "-0x8000000000000000000000000000000000000000000000000000000000000000"
	int256ModString = "0x8000000000000000000000000000000000000000000000000000000000000000"
)

var (
	int256Min = new(big.Int)
	int256Mod = new(big.Int)
)

func init() {
	int256Min.SetString(int256MinString, 0)
	int256Mod.SetString(int256ModString, 0)
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
