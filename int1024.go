package hash

import (
	"math/big"
)

var (
	int1024Min = new(big.Int)
	int1024Mod = new(big.Int)
)

func init() {
	int1024Mod.Lsh(bigOne, 1023)
	int1024Min.Neg(int1024Mod)
}

func Int1024(data []byte) (hash *big.Int) {
	hash = Uint1024(data)

	delta := new(big.Int)
	delta.Sub(hash, int1024Mod)

	if delta.Cmp(bigZero) >= 0 {
		hash.Add(int1024Min, delta)
	}

	return
}

func Int1024String(data string) (hash *big.Int) {
	hash = Int1024([]byte(data))

	return
}
