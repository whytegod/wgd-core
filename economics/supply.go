package economics

import "github.com/whytegod/wgd-core/genesis"

func CalculateBlockReward(height int, cfg genesis.Config) float64 {
	halvings := height / cfg.HalvingInterval

	reward := cfg.BlockReward

	for i := 0; i < halvings; i++ {
		reward = reward / 2
	}

	if reward < 0.00000001 {
		return 0
	}

	return reward
}