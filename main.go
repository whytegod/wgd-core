package main

import (
	"fmt"
	"wgd-core/core"
)

func main() {
	bc := core.NewBlockchain()

	bc.AddBlock("First Block After Genesis")
	bc.AddBlock("Second Block After Genesis")

	for _, block := range bc.Blocks {
		fmt.Println("Index:", block.Index)
		fmt.Println("Timestamp:", block.Timestamp)
		fmt.Println("Data:", block.Data)
		fmt.Println("PrevHash:", block.PrevHash)
		fmt.Println("Hash:", block.Hash)
		fmt.Println("----------------------------------")
	}
}