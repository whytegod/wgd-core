package main

import (
	"fmt"
	"wgd-core/ledger"
)

func main() {

	chain := ledger.NewBlockchain()

	fmt.Println("====================================")
	fmt.Println("Whytegod Blockchain Started")
	fmt.Println("Current Supply:", chain.TotalSupply, "WGD")
	fmt.Println("Blocks:", len(chain.Blocks))
	fmt.Println("====================================")

	chain.MineBlock("First real mined block")

	fmt.Println("After Mining:")
	fmt.Println("Current Supply:", chain.TotalSupply, "WGD")
	fmt.Println("Blocks:", len(chain.Blocks))
}