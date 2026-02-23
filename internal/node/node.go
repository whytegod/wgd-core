package node

import (
	"sync"

	"wgd-core/internal/block"
	"wgd-core/internal/ledger"
)

type Node struct {
	mu      sync.RWMutex
	Chain   []*block.Block
	UTXO    *ledger.UTXOSet
	Mempool map[[32]byte]*block.Transaction
}

func NewNode() *Node {
	return &Node{
		Chain:   make([]*block.Block, 0),
		UTXO:    ledger.NewUTXOSet(),
		Mempool: make(map[[32]byte]*block.Transaction),
	}
}

func (n *Node) AddBlock(b *block.Block) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.Chain = append(n.Chain, b)
}