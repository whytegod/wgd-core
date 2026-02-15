package main

import (
	"fmt"

	"wgd-core/ledger"
)

const ProtocolName = "Whytegod"
const ProtocolVersion = "v0.1.0"

func main() {

	fmt.Println("===================================")
	fmt.Println("Starting", ProtocolName, ProtocolVersion)
	fmt.Println("===================================")

	chain := ledger.NewLedger(9720000)

	fmt.Println("Genesis Block Created")
	fmt.Println("Total Supply:", chain.TotalSupply())
	fmt.Println("Block Height:", 1)
}