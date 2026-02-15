package main

import (
	"fmt"
	"wgd-core/ledger"
)

const (
	ProtocolName    = "Whytegod"
	ProtocolVersion = "v0.3.0"
	InitialSupply   = 9720000.0
	BlockReward     = 10.0
)

func main() {

	fmt.Println("====================================")
	fmt.Println("Starting", ProtocolName, ProtocolVersion)
	fmt.Println("====================================")

	// Initialize Ledger
	l := ledger.NewLedger(InitialSupply, BlockReward)

	fmt.Println("Genesis Height:", l.Height())
	fmt.Println("Total Supply:", l.TotalSupply())
	fmt.Println("Treasury Balance:", l.TreasuryBalance())

	// Mine Some Blocks
	l.AddBlock("First Platinum Block")
	l.AddBlock("Second Platinum Block")

	fmt.Println("\nAfter Mining:")
	fmt.Println("Height:", l.Height())
	fmt.Println("Total Supply:", l.TotalSupply())
	fmt.Println("Treasury Balance:", l.TreasuryBalance())

	fmt.Println("\nFull Chain:")
	l.PrintChain()
}