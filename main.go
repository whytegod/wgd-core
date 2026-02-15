package main

import (
	"fmt"

	"github.com/whytegod/wgd-core/genesis"
)

const ProtocolName = "Whytegod"
const ProtocolVersion = "v0.1.0"

func main() {

	cfg := genesis.DefaultGenesis()

	fmt.Println("===================================")
	fmt.Println("Starting", ProtocolName, ProtocolVersion)
	fmt.Printf("Initial Supply: %.0f WGD\n", cfg.InitialSupply)
	fmt.Printf("Supply Cap: %.0f WGD\n", cfg.AsymptoticSupplyCap)
	fmt.Println("===================================")
}