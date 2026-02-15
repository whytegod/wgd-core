package main

import (
	"fmt"

	"github.com/whytegod/wgd-core/core"
)

func main() {

	// Create blockchain
	chain := core.NewBlockchain()

	// Create wallets
	walletA := core.NewWallet()
	walletB := core.NewWallet()

	// Create transaction
	tx := core.Transaction{
		From:   walletA.PublicKey,
		To:     walletB.PublicKey,
		Amount: 10,
	}

	tx.Sign(walletA.PrivateKey)

	// Add transaction to mempool
	chain.AddTransaction(tx)

	// Mine block
	chain.MineBlock(walletA.PublicKey)

	// Print blockchain
	for _, block := range chain.Blocks {
		fmt.Println("=================================")
		fmt.Println("Block:", block.Index)
		fmt.Println("Hash:", block.Hash)
		fmt.Println("PrevHash:", block.PrevHash)
		fmt.Println("Transactions:", block.Transactions)
	}

	// Save chain
	chain.SaveToFile("whytegod_chain.json")
}