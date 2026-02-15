package main

import (
	"fmt"
	"github.com/YOUR_GITHUB_USERNAME/whytegod-core/genesis"
)

const ProtocolName = "Whytegod"
const ProtocolVersion = "v0.1.0"

func main() {
	cfg := genesis.DefaultGenesis()

	fmt.Println("Starting", ProtocolName, ProtocolVersion)
	fmt.Println("Initial Supply:", cfg.InitialSupply)
	fmt.Println("Supply Cap:", cfg.AsymptoticSupplyCap)
}