package node

import (
	"sync"
	"wgd-core/internal/block"
	"wgd-core/internal/ledger"
)

type Node struct {
	mu       sync.RWMutex
	Chain    []*block.BlockHeader
	UTXO     *ledger.UTXOSet
	Mempool  map[[32]byte]interface{}
}

func NewNode() *Node {
	return &Node{
		Chain:   make([]*block.BlockHeader, 0),
		UTXO:    ledger.NewUTXOSet(),
		Mempool: make(map[[32]byte]interface{}),
	}
}