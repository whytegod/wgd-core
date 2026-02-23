package consensus

const (
	InitialRewardWerts = 4625000000
	HalvingInterval    = 105120
	MaxSupplyWerts     = 9720000 * 100000000
)

func BlockReward(height uint64) uint64 {
	halvings := height / HalvingInterval

	reward := InitialRewardWerts >> halvings

	return reward
}