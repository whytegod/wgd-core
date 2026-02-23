package main

import (
	"wgd-core/internal/core"
)

func main() {
	node := core.NewNode()

	node.AddBlock([]string{"tx1", "tx2"})
	node.AddBlock([]string{"tx3"})

	node.PrintState()
}