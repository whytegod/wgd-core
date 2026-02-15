package main

import (
	"fmt"

	"github.com/whytegod/wgd-core/genesis"
	"github.com/whytegod/wgd-core/ledger"
)

const ProtocolName = "Whytegod"
const ProtocolVersion = "v0.3.0"

func main() {

	cfg := genesis.DefaultGenesis()

	// Initialize Blockchain Ledger
	chain := ledger.NewLedger(
		cfg.InitialSupply,
		50, // Block reward
	)

	fmt.Println("===================================")
	fmt.Println("Starting", ProtocolName, ProtocolVersion)
	fmt.Println("Genesis Height:", chain.Height())
	fmt.Printf("Total Supply: %.0f WGD\n", chain.TotalSupply())
	fmt.Printf("Treasury Balance: %.0f WGD\n", chain.TreasuryBalance())
	fmt.Println("===================================")

	// Mine Example Blocks
	chain.AddBlock("First Whytegod Block")
	chain.AddBlock("Second Whytegod Block")

	fmt.Println("\nAfter Mining:")
	fmt.Println("Height:", chain.Height())
	fmt.Printf("Total Supply: %.0f WGD\n", chain.TotalSupply())
	fmt.Printf("Treasury Balance: %.0f WGD\n", chain.TreasuryBalance())

	fmt.Println("\nFull Chain:")
	chain.PrintChain()
}