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
		Amount: 10,
	}

	bc.AddTransaction(tx)
	bc.MineBlock("Miner1")

	for _, block := range bc.Blocks {
		fmt.Println("Block:", block.Index)
		fmt.Println("Prev Hash:", block.PrevHash)
		fmt.Println("Hash:", block.Hash)
		fmt.Println("Transactions:", block.Transactions)
		fmt.Println("--------------------")
	}
}