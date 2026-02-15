package main

import (
	"fmt"

	"github.com/whytegod/wgd-core/genesis"
)

const (
	ProtocolName    = "Whytegod"
	ProtocolVersion = "v0.1.0"
)

func main() {
	startProtocol()
}

func startProtocol() {
	config := genesis.DefaultGenesis()

	printBanner()
	printProtocolInfo(config)
}

func printBanner() {
	fmt.Println("===================================")
}

func printProtocolInfo(cfg genesis.GenesisConfig) {
	fmt.Printf("Starting %s %s\n", ProtocolName, ProtocolVersion)
	fmt.Printf("Initial Supply: %.0f WGD\n", cfg.InitialSupply)
	fmt.Printf("Supply Cap: %.0f WGD\n", cfg.AsymptoticSupplyCap)
	fmt.Println("===================================")
}