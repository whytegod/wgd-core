package genesis

type GenesisConfig struct {
	InitialSupply        uint64
	AsymptoticSupplyCap  uint64
	InitialBlockReward   uint64
	DecayConstant        float64
	MicroTailEmission    uint64
	MinValidatorStake    uint64
	EpochLength          uint64
	CheckpointInterval   uint64
}

func DefaultGenesis() GenesisConfig {
	return GenesisConfig{
		InitialSupply:       1_000_000_000,
		AsymptoticSupplyCap: 10_000_000_000,
		InitialBlockReward:  50,
		DecayConstant:       0.0001,
		MicroTailEmission:   1,
		MinValidatorStake:   10_000,
		EpochLength:         100,
		CheckpointInterval:  500,
	}
}