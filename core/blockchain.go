package core

import (
	"errors"
	"fmt"
)

// Blockchain is the in-memory representation of the chain.
type Blockchain struct {
	Blocks     []*Block
	Mempool    []*Transaction
	Difficulty int
}

// NewBlockchain creates a chain with a genesis block.
func NewBlockchain() *Blockchain {
	genesis := CreateGenesisBlock()
	return &Blockchain{
		Blocks:     []*Block{genesis},
		Mempool:    []*Transaction{},
		Difficulty: 2, // default difficulty (tweakable)
	}
}

// CreateGenesisBlock builds the genesis block. Adjust coinbase amount as needed.
func CreateGenesisBlock() *Block {
	coinbase := NewCoinbaseTx("genesis-address", 50)
	txs := []*Transaction{coinbase}
	return NewBlock(0, txs, []byte{}, 2)
}

// AddTransaction places a transaction into mempool.
func (bc *Blockchain) AddTransaction(tx *Transaction) {
	if tx == nil {
		return
	}
	bc.Mempool = append(bc.Mempool, tx)
}

// AddBlock appends a block after basic validation.
func (bc *Blockchain) AddBlock(b *Block) error {
	if b == nil {
		return errors.New("nil block")
	}
	last := bc.Blocks[len(bc.Blocks)-1]
	// ensure index and prevhash match
	if b.Index != last.Index+1 {
		return fmt.Errorf("invalid index: want %d got %d", last.Index+1, b.Index)
	}
	if !bytesEqual(b.PrevHash, last.Hash) {
		return fmt.Errorf("prev hash mismatch")
	}
	// append and done
	bc.Blocks = append(bc.Blocks, b)
	return nil
}

// MineBlock collects mempool txs, creates coinbase, mines block, appends to chain and clears mempool.
func (bc *Blockchain) MineBlock(minerAddress string) (*Block, error) {
	if minerAddress == "" {
		return nil, fmt.Errorf("miner address required")
	}
	coinbase := NewCoinbaseTx(minerAddress, 50)
	txs := []*Transaction{coinbase}
	txs = append(txs, bc.Mempool...)

	prev := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(prev.Index+1, txs, prev.Hash, bc.Difficulty)

	if err := bc.AddBlock(newBlock); err != nil {
		return nil, err
	}

	// clear mempool (in real chain you'd remove only included txs & handle reorgs)
	bc.Mempool = []*Transaction{}
	return newBlock, nil
}

// bytesEqual helper
func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}