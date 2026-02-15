package main

import (
	"fmt"
	"github.com/YOUR_GITHUB_USERNAME/whytegod-core/genesis"
)

func main() {
	cfg := genesis.DefaultGenesis()

	fmt.Println("Starting", cfg.ProtocolName, cfg.ProtocolVersion)
	fmt.Println("Initial Supply:", cfg.InitialSupply)
	fmt.Println("Supply Cap:", cfg.AsymptoticSupplyCap)
	fmt.Println("Block Time:", cfg.BlockTimeSeconds, "seconds")
	fmt.Println("Max Block Size:", cfg.MaxBlockSizeBytes, "bytes")
}