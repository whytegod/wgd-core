package main

import (
	"fmt"

	"github.com/whytegod/wgd-core/core"
)

func main() {

	bc := core.NewBlockchain()

	tx := core.Transaction{
		From:   "Alice",
		To:     "Bob",
		Amount: 50,
	}

	bc.AddTransaction(tx)
	bc.MineBlock("WhytegodMiner")

	for _, block := range bc.Blocks {
		fmt.Println("Index:", block.Index)
		fmt.Println("Previous Hash:", block.PrevHash)
		fmt.Println("Hash:", block.Hash)
		fmt.Println("Transactions:", block.Transactions)
		fmt.Println("-------------------------")
	}
}