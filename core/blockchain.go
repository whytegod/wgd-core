package core

import (
	"github.com/whytegod/wgd-core/block"
)

type Blockchain struct {
	Blocks []block.Block
}

// Create a new blockchain with genesis block
func NewBlockchain() *Blockchain {
	genesis := block.NewBlock(
		0,
		"0",
		"Genesis Block - Whytegod",
		50.0,
	)

	return &Blockchain{
		Blocks: []block.Block{genesis},
	}
}

// Add a new block
func (bc *Blockchain) AddBlock(data string, reward float64) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]

	newBlock := block.NewBlock(
		prevBlock.Index+1,
		prevBlock.Hash,
		data,
		reward,
	)

	bc.Blocks = append(bc.Blocks, newBlock)
}