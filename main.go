package main

import (
	"fmt"
	"wgd-core/core"
)

func main() {
	// create blockchain
	bc := core.NewBlockchain()

	// simulate adding a normal transaction (optional)
	tx := core.NewCoinbaseTx("alice", 10)
	bc.AddTransaction(tx)

	// mine block
	block, err := bc.MineBlock("miner1")
	if err != nil {
		panic(err)
	}

	fmt.Println("Block mined!")
	fmt.Printf("Index: %d\n", block.Index)
	fmt.Printf("Hash: %x\n", block.Hash)
	fmt.Printf("PrevHash: %x\n", block.PrevHash)
	fmt.Printf("Transactions: %d\n", len(block.Transactions))
}