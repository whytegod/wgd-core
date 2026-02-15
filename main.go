package main

import (
	"fmt"

	"github.com/whytegod/wgd-core/core"
	"github.com/whytegod/wgd-core/genesis"
)

const ProtocolName = "Whytegod"
const ProtocolVersion = "v0.1.0"

func main() {

	// Load Genesis Configuration
	cfg := genesis.DefaultGenesis()

	fmt.Println("===================================")
	fmt.Println("Starting", ProtocolName, ProtocolVersion)
	fmt.Println("Initial Supply:", cfg.InitialSupply, "WGD")
	fmt.Println("Supply Cap:", cfg.AsymptoticSupplyCap, "WGD")
	fmt.Println("===================================")

	// Initialize Blockchain
	chain := core.NewBlockchain()

	// Add Some Blocks
	chain.AddBlock("First transaction", 50.0)
	chain.AddBlock("Second transaction", 45.0)
	chain.AddBlock("Third transaction", 40.0)

	// Print Blockchain
	for _, block := range chain.Blocks {
		fmt.Println("---------------")
		fmt.Println("Index:", block.Index)
		fmt.Println("Timestamp:", block.Timestamp)
		fmt.Println("Data:", block.Data)
		fmt.Println("PrevHash:", block.PrevHash)
		fmt.Println("Hash:", block.Hash)
		fmt.Println("Reward:", block.Reward)
	}
}