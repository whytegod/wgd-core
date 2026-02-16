package main

import (
	"fmt"
	"wgd-core/core"
)

func main() {
	bc := core.NewBlockchain()

	bc.AddBlock("First Block")
	bc.AddBlock("Second Block")

	for _, block := range bc.Blocks {
		fmt.Println("Index:", block.Index)
		fmt.Println("Timestamp:", block.Timestamp)
		fmt.Println("Data:", block.Data)
		fmt.Println("Prev Hash:", block.PrevHash)
		fmt.Println("Hash:", block.Hash)
		fmt.Println("----------------------------")
	}
}