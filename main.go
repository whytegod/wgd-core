package main

import (
	"fmt"

	"github.com/whytegod/wgd-core/genesis"
)

const (
	ProtocolName    = "Whytegod"
	ProtocolVersion = "v0.1.0"
)

func main() {
	// Load Genesis Configuration
	cfg := genesis.DefaultGenesis()

	fmt.Println("===================================")
	fmt.Println("Starting", ProtocolName, ProtocolVersion)
	fmt.Println("===================================")
	fmt.Println("Initial Supply:", cfg.InitialSupply)
	fmt.Println("Asymptotic Supply Cap:", cfg.AsymptoticSupplyCap)
	fmt.Println("Block Reward:", cfg.BlockReward)
	fmt.Println("Halving Interval:", cfg.HalvingInterval)
	fmt.Println("===================================")
}