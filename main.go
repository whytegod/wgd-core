package main

import (
	"fmt"

	"github.com/whytegod/wgd-core/core"
	"github.com/whytegod/wgd-core/genesis"
	"github.com/whytegod/wgd-core/ledger"
)

const (
	ProtocolName    = "Whytegod"
	ProtocolVersion = "v0.2.0"
)

func main() {

	// ----------------------------------
	// 1️⃣ Load Genesis Configuration
	// ----------------------------------
	cfg := genesis.DefaultGenesis()

	// ----------------------------------
	// 2️⃣ Initialize Ledger
	// ----------------------------------
	ldg := ledger.NewLedger(cfg.InitialSupply)

	// ----------------------------------
	// 3️⃣ Initialize Blockchain
	// ----------------------------------
	chain := core.NewBlockchain()

	// ----------------------------------
	// 4️⃣ Boot Information
	// ----------------------------------
	fmt.Println("===================================")
	fmt.Println("Protocol:", ProtocolName)
	fmt.Println("Version :", ProtocolVersion)
	fmt.Println("Initial Supply:", ldg.TotalSupply(), "WGD")
	fmt.Println("Treasury Balance:", ldg.BalanceOf("treasury"), "WGD")
	fmt.Println("===================================")

	// ----------------------------------
	// 5️⃣ Create Transaction
	// ----------------------------------
	tx1 := core.Transaction{
		From:   "treasury",
		To:     "alice",
		Amount: 1000,
	}

	// ----------------------------------
	// 6️⃣ Apply Transaction to Ledger
	// ----------------------------------
	err := ldg.Transfer(tx1.From, tx1.To, tx1.Amount)
	if err != nil {
		fmt.Println("Transaction failed:", err)
		return
	}

	// ----------------------------------
	// 7️⃣ Add Transaction to Blockchain
	// ----------------------------------
	chain.AddBlock([]core.Transaction{tx1})

	// ----------------------------------
	// 8️⃣ Display Updated Ledger State
	// ----------------------------------
	fmt.Println("\n--- Ledger State ---")
	fmt.Println("Treasury:", ldg.BalanceOf("treasury"), "WGD")
	fmt.Println("Alice   :", ldg.BalanceOf("alice"), "WGD")

	// ----------------------------------
	// 9️⃣ Display Blockchain
	// ----------------------------------
	fmt.Println("\n--- Blockchain ---")
	for _, block := range chain.Blocks {
		fmt.Println("Block Index:", block.Index)
		fmt.Println("Timestamp  :", block.Timestamp)
		fmt.Println("Hash       :", block.Hash)
		fmt.Println("PrevHash   :", block.PrevHash)
		fmt.Println("Tx Count   :", len(block.Transactions))
		fmt.Println("-----------------------------")
	}

	fmt.Println("\nNode shutdown complete.")
}