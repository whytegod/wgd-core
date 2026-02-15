package genesis

type Config struct {
	InitialSupply        float64
	AsymptoticSupplyCap  float64
	BlockReward          float64
	HalvingInterval      int
}

func DefaultGenesis() Config {
	return Config{
		InitialSupply:       0,
		AsymptoticSupplyCap: 21000000,
		BlockReward:         50,
		HalvingInterval:     210000,
	}
}