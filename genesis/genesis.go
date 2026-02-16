cat > genesis/genesis.go << 'EOF'
package genesis

type Config struct {
	ProtocolName    string
	ProtocolVersion string
	InitialSupply   uint64
	SupplyCap       uint64
}

func DefaultGenesis() Config {
	return Config{
		ProtocolName:    "Whytegod",
		ProtocolVersion: "v0.1.0",
		InitialSupply:   9720000,
		SupplyCap:       9720000,
	}
}
EOF