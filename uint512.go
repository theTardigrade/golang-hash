package hash

import (
	"math/big"
)

const (
	uint512PrimeString  = "0x100000000000000000000000000000000000000000000000000000000000000000000000000000000000157"
	uint512OffsetString = "0xb86db0b1171f4416dca1e50f309990acac87d059c90000000000000000000d21e948f68a34c192f62ea79bc942dbe7c182036415f56e34bac982aac4afe9fd9"
)

var (
	uint512Mod    = new(big.Int)
	uint512Prime  = new(big.Int)
	uint512Offset = new(big.Int)
)

func init() {
	uint512Mod.Lsh(bigOne, 512)
	uint512Prime.SetString(uint512PrimeString, 0)
	uint512Offset.SetString(uint512OffsetString, 0)
}

func Uint512(data []byte) (hash *big.Int) {
	hash = new(big.Int)
	hash.Set(uint512Offset)

	datum := new(big.Int)
	for _, d := range data {
		hash.Xor(hash, datum.SetUint64(uint64(d)))
		hash.Mul(hash, uint512Prime)
		hash.Mod(hash, uint512Mod)
	}

	return
}

func Uint512String(data string) (hash *big.Int) {
	hash = Uint512([]byte(data))

	return
}
