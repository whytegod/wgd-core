package block

import "math/big"

func BlockWork(bits uint32) *big.Int {
	target := CompactToTarget(bits)
	maxTarget := new(big.Int).Lsh(big.NewInt(1), 256)
	work := new(big.Int).Div(maxTarget, target)
	return work
}