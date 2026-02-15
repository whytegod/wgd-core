package main

import (
	"fmt"

	"github.com/whytegod/wgd-core/ledger"
)

const ProtocolName = "Whytegod"
const ProtocolVersion = "v0.2.0"

func main() {

	fmt.Println("===================================")
	fmt.Println("Starting", ProtocolName, ProtocolVersion)
	fmt.Println("===================================")

	chain := ledger.NewLedger(9720000, 50)

	fmt.Println("Genesis Loaded")
	fmt.Println("Height:", chain.Height())
	fmt.Println("Supply:", chain.TotalSupply())
}