package main

import (
	"fmt"
	"log"
	"time"

	"wgd-core/internal/block"
	"wgd-core/internal/core"
	"wgd-core/internal/economics"
	"wgd-core/internal/genesis"
	"wgd-core/internal/ledger"
)

func main() {

	fmt.Println("====================================")
	fmt.Println("        WGD CORE — Digital Platinum")
	fmt.Println("====================================")

	//-----------------------------------
	// 1️⃣ Initialize Ledger
	//-----------------------------------
	ledgerState := ledger.NewUTXOSet()
	if ledgerState == nil {
		log.Fatal("Failed to initialize ledger")
	}

	//-----------------------------------
	// 2️⃣ Apply Genesis Block
	//-----------------------------------
	genesisConfig := genesis.DefaultGenesis()

	err := genesis.ApplyGenesis(ledgerState, genesisConfig)
	if err != nil {
		log.Fatalf("Genesis failed: %v", err)
	}

	fmt.Println("Genesis applied successfully")

	//-----------------------------------
	// 3️⃣ Initialize Economics Engine
	//-----------------------------------
	econConfig := economics.Config{
		MaxSupply:       9_720_000,
		BlockReward:     50,
		HalvingInterval: 730,
		InitialSupply:   genesisConfig.InitialSupply,
	}

	econ := economics.NewEconomicsEngine(econConfig)
	if econ == nil {
		log.Fatal("Failed to initialize economics engine")
	}

	//-----------------------------------
	// 4️⃣ Initialize Blockchain Core
	//-----------------------------------
	chain := core.NewBlockchain(ledgerState, econ)
	if chain == nil {
		log.Fatal("Failed to initialize blockchain")
	}

	fmt.Println("Blockchain initialized")

	//-----------------------------------
	// 5️⃣ Simulate Mining 3 Blocks
	//-----------------------------------
	for i := 0; i < 3; i++ {

		fmt.Printf("\nMining block %d...\n", i+1)

		prevHash := chain.LastBlockHash()

		newBlock, err := block.CreateBlock(
			prevHash,
			nil,
			time.Now(),
		)
		if err != nil {
			log.Fatalf("Block creation failed: %v", err)
		}

		err = chain.AddBlock(newBlock)
		if err != nil {
			log.Fatalf("Failed to add block: %v", err)
		}

		reward := econ.CalculateReward(chain.Height())

		err = ledgerState.ApplyMiningReward(reward)
		if err != nil {
			log.Fatalf("Failed to apply reward: %v", err)
		}

		fmt.Printf("Block %d added successfully\n", newBlock.Height)
		fmt.Printf("Current Supply: %d\n", econ.CurrentSupply())
	}

	//-----------------------------------
	// 6️⃣ Final Chain Status
	//-----------------------------------
	fmt.Println("\n====================================")
	fmt.Println("Blockchain Height:", chain.Height())
	fmt.Printf("Total Supply: %d\n", econ.CurrentSupply())
	fmt.Println("System running correctly ✅")
	fmt.Println("====================================")
}