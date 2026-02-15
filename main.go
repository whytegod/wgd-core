package main

import (
	"fmt"
)

const ProtocolName = "Whytegod"
const ProtocolVersion = "v0.1.0"

func main() {

	printBanner()

	fmt.Println("Node initialized successfully.")
}

func printBanner() {
	fmt.Println("======================================")
	fmt.Printf("  %s %s\n", ProtocolName, ProtocolVersion)
	fmt.Println("======================================")
}