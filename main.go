package main

import (
	"fmt"
	"log"
	"time"

	"github.com/whytegod/wgd-core/block"
	"github.com/whytegod/wgd-core/core"
	"github.com/whytegod/wgd-core/economics"
	"github.com/whytegod/wgd-core/genesis"
	"github.com/whytegod/wgd-core/ledger"
)

func main() {

	fmt.Println("======================================")
	fmt.Println("      WGD CORE BLOCKCHAIN NODE       ")
	fmt.Println("======================================")

	// --------------------------------------------------
	// 1. Initialize Ledger (State)
	// --------------------------------------------------
	state := ledger.NewLedger()

	fmt.Println("[✓] Ledger initialized")

	// --------------------------------------------------
	// 2. Load Genesis Configuration
	// --------------------------------------------------
	genesisConfig := genesis.DefaultGenesis()

	err := genesis.ApplyGenesis(state, genesisConfig)
	if err != nil {
		log.Fatalf("Genesis initialization failed: %v", err)
	}

	fmt.Println("[✓] Genesis block applied")

	// --------------------------------------------------
	// 3. Initialize Economics Engine
	// --------------------------------------------------
	econ := economics.NewEconomicsEngine(
		economics.Config{
			MaxSupply:        21_000_000,
			BlockReward:      50,
			HalvingInterval:  100,
			InitialSupply:    genesisConfig.InitialSupply,
		},
	)

	fmt.Println("[✓] Economics engine loaded")

	// --------------------------------------------------
	// 4. Initialize Blockchain Core
	// --------------------------------------------------
	chain := core.NewBlockchain(state, econ)

	fmt.Println("[✓] Blockchain core started")

	// --------------------------------------------------
	// 5. Simulate Block Production
	// --------------------------------------------------

	for i := 0; i < 5; i++ {

		fmt.Printf("\n⛏ Mining block %d...\n", i+1)

		newBlock, err := block.CreateBlock(
			chain.LastBlockHash(),
			state.PendingTransactions(),
			time.Now(),
		)

		if err != nil {
			log.Fatalf("Block creation failed: %v", err)
		}

		err = chain.AddBlock(newBlock)
		if err != nil {
			log.Fatalf("Block rejected: %v", err)
		}

		fmt.Printf("Block %d added successfully\n", newBlock.Height)
		fmt.Printf("Current Supply: %d\n", econ.CurrentSupply())
	}

	// --------------------------------------------------
	// 6. Final Chain Summary
	// --------------------------------------------------

	fmt.Println("\n======================================")
	fmt.Println("CHAIN SUMMARY")
	fmt.Println("======================================")

	fmt.Printf("Total Blocks: %d\n", chain.Height())
	fmt.Printf("Final Supply: %d\n", econ.CurrentSupply())

	fmt.Println("Node shutting down cleanly...")
}