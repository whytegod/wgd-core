package economics

import "math"

type SupplyModel struct {
	InitialSupply float64
	Cap           float64
	GrowthRate    float64
}

func (s SupplyModel) SupplyAtBlock(blockHeight float64) float64 {
	return s.Cap - (s.Cap-s.InitialSupply)*math.Exp(-s.GrowthRate*blockHeight)
}