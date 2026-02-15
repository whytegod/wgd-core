package main

import (
	"fmt"

	"github.com/whytegod/wgd-core/genesis"
	"github.com/whytegod/wgd-core/ledger"
)

const ProtocolName = "Whytegod"
const ProtocolVersion = "v0.1.0"

func main() {
	cfg := genesis.DefaultGenesis()

	ldg := ledger.NewLedger(cfg.InitialSupply)

	fmt.Println("===================================")
	fmt.Println("Starting", ProtocolName, ProtocolVersion)
	fmt.Println("Initial Supply:", ldg.TotalSupply(), "WGD")
	fmt.Println("Treasury Balance:", ldg.BalanceOf("treasury"), "WGD")
	fmt.Println("===================================")

	// Test transfer
	err := ldg.Transfer("treasury", "alice", 1000)
	if err != nil {
		fmt.Println("Transfer failed:", err)
		return
	}

	fmt.Println("After Transfer:")
	fmt.Println("Treasury:", ldg.BalanceOf("treasury"))
	fmt.Println("Alice:", ldg.BalanceOf("alice"))
}