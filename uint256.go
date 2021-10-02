package hash

import (
	"math/big"
)

const (
	uint256PrimeHexString  = "1000000000000000000000000000000000000000163"
	uint256OffsetHexString = "dd268dbcaac550362d98c384c4e576ccc8b1536847b6bbb31023b4c8caee0535"
)

var (
	uint256Mod    = new(big.Int)
	uint256Prime  = new(big.Int)
	uint256Offset = new(big.Int)
)

func init() {
	uint256Mod.Lsh(bigOne, 256)
	uint256Prime.SetString(uint256PrimeHexString, 16)
	uint256Offset.SetString(uint256OffsetHexString, 16)
}

func Uint256(data []byte) (hash *big.Int) {
	hash = new(big.Int)
	hash.Set(uint256Offset)

	datum := new(big.Int)
	for _, d := range data {
		hash.Xor(hash, datum.SetUint64(uint64(d)))
		hash.Mul(hash, uint256Prime)
		hash.Mod(hash, uint256Mod)
	}

	return
}

func Uint256String(data string) (hash *big.Int) {
	hash = Uint256([]byte(data))

	return
}
