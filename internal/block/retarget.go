package block

import (
	"math/big"
)

const (
	TargetBlockTime = 600
	AdjustmentWindow = 10
	MaxAdjustUpFactor = 2
	MaxAdjustDownFactor = 2
)

func CalculateNewBits(prevBits uint32, timestamps []uint32) uint32 {

	if len(timestamps) < 2 {
		return prevBits
	}

	first := timestamps[0]
	last := timestamps[len(timestamps)-1]

	actualTime := int64(last - first)
	expectedTime := int64(TargetBlockTime * (len(timestamps)-1))

	if actualTime <= 0 {
		actualTime = 1
	}

	prevTarget := CompactToTarget(prevBits)

	newTarget := new(big.Int).Mul(prevTarget, big.NewInt(actualTime))
	newTarget.Div(newTarget, big.NewInt(expectedTime))

	maxUp := new(big.Int).Mul(prevTarget, big.NewInt(MaxAdjustUpFactor))
	maxDown := new(big.Int).Div(prevTarget, big.NewInt(MaxAdjustDownFactor))

	if newTarget.Cmp(maxUp) > 0 {
		newTarget = maxUp
	}

	if newTarget.Cmp(maxDown) < 0 {
		newTarget = maxDown
	}

	return TargetToCompact(newTarget)
}