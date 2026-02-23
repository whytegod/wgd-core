package node

import (
	"errors"
	"wgd-core/internal/block"
	"wgd-core/internal/consensus"
)

func (n *Node) AddBlock(b *block.Block) error {

	n.mu.Lock()
	defer n.mu.Unlock()

	prevNode, exists := n.Blocks[b.Header.PrevHash]
	if !exists {
		return errors.New("parent block not found")
	}

	// Validate block
	if err := consensus.ValidateBlock(b, prevNode.Block, n.UTXO); err != nil {
		return err
	}

	newWork := block.BlockWork(b.Header.Bits)

	totalWork := new(big.Int).Add(prevNode.TotalWork, newWork)

	newNode := &BlockNode{
		Block:     b,
		Parent:    prevNode,
		TotalWork: totalWork,
	}

	hash := b.Header.Hash()
	n.Blocks[hash] = newNode

	// Check if new chain is stronger
	if n.BestTip == nil || totalWork.Cmp(n.BestTip.TotalWork) > 0 {
		n.reorganize(newNode)
	}

	return nil
}