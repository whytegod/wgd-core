package main

import (
	"fmt"
	"time"

	"github.com/whytegod/wgd-core/block"
	"github.com/whytegod/wgd-core/economics"
	"github.com/whytegod/wgd-core/genesis"
	"github.com/whytegod/wgd-core/ledger"
)

const (
	ProtocolName    = "Whytegod"
	ProtocolVersion = "v0.1.0"
)

func main() {

	fmt.Println("===================================")
	fmt.Println("Launching", ProtocolName, ProtocolVersion)
	fmt.Println("Digital Platinum Standard")
	fmt.Println("===================================")

	// Load Genesis
	cfg := genesis.DefaultGenesis()

	// Initialize Ledger
	chain := ledger.NewLedger()

	// Create Genesis Block
	genBlock := block.NewBlock(0, "0", "Genesis Block", cfg.BlockReward)

	chain.AddBlock(genBlock)

	fmt.Println("Genesis Block Created")
	fmt.Println("Initial Supply:", cfg.InitialSupply)
	fmt.Println("Supply Cap:", cfg.AsymptoticSupplyCap)

	// Simulate next block
	reward := economics.CalculateBlockReward(1, cfg)
	newBlock := block.NewBlock(1, genBlock.Hash, "Second Block", reward)

	chain.AddBlock(newBlock)

	fmt.Println("Second Block Added")
	fmt.Println("Current Supply:", chain.TotalSupply())
	fmt.Println("Blockchain Height:", chain.Height())
	fmt.Println("Timestamp:", time.Now())
	fmt.Println("===================================")
}