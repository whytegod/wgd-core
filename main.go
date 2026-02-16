package main

import (
	"fmt"
	"wgd-core/core"
)

func main() {
	bc := core.NewBlockchain()

	priv1, addr1 := core.GenerateKeyPair()
	_, addr2 := core.GenerateKeyPair()

	tx := core.Transaction{
		From:   addr1,
		To:     addr2,
		Amount: 10,
	}

	tx.Hash = tx.CalculateHash()
	tx.SignTransaction(priv1)

	bc.AddTransaction(tx)

	fmt.Println("⛏ Mining block...")
	bc.MinePendingTransactions()

	for _, block := range bc.Blocks {
		fmt.Println("Index:", block.Index)
		fmt.Println("Hash:", block.Hash)
		fmt.Println("Prev:", block.PrevHash)
		fmt.Println("Nonce:", block.Nonce)
		fmt.Println("Tx Count:", len(block.Transactions))
		fmt.Println("------------------------")
	}
}