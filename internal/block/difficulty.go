package block

import (
	"math/big"
)

func CompactToTarget(bits uint32) *big.Int {
	exponent := bits >> 24
	coefficient := bits & 0x007fffff

	target := new(big.Int).SetUint64(uint64(coefficient))
	shift := 8 * (exponent - 3)

	target.Lsh(target, uint(shift))
	return target
}

func TargetToCompact(target *big.Int) uint32 {
	bytes := target.Bytes()
	size := len(bytes)

	var compact uint32
	if size <= 3 {
		value := new(big.Int).Set(target)
		value.Lsh(value, uint(8*(3-size)))
		compact = uint32(value.Uint64())
	} else {
		value := new(big.Int).Rsh(target, uint(8*(size-3)))
		compact = uint32(value.Uint64())
	}

	if compact&0x00800000 != 0 {
		compact >>= 8
		size++
	}

	compact |= uint32(size) << 24
	return compact
}