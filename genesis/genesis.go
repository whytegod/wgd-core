package genesis

// GenesisConfig defines the initial monetary configuration
type GenesisConfig struct {
	InitialSupply       float64
	AsymptoticSupplyCap float64
}

// DefaultGenesis returns Whytegod's platinum standard genesis configuration
func DefaultGenesis() GenesisConfig {
	return GenesisConfig{
		InitialSupply:       9720000.0,  // 9.72 million WGD initial issuance
		AsymptoticSupplyCap: 12000000.0, // 12 million asymptotic cap
	}
}