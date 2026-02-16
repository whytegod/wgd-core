package core

import (
	"fmt"
)

type Blockchain struct {
	Blocks     []*Block
	Mempool    []*Transaction
	Difficulty int
}

func NewBlockchain() *Blockchain {
	genesis := CreateGenesisBlock()

	return &Blockchain{
		Blocks:     []*Block{genesis},
		Mempool:    []*Transaction{},
		Difficulty: 2,
	}
}

func CreateGenesisBlock() *Block {
	coinbase := NewCoinbaseTx("genesis-address", 50)
	return NewBlock(0, []*Transaction{coinbase}, []byte{}, 2)
}

func (bc *Blockchain) AddTransaction(tx *Transaction) {
	bc.Mempool = append(bc.Mempool, tx)
}

func (bc *Blockchain) MineBlock(minerAddress string) {
	coinbase := NewCoinbaseTx(minerAddress, 50)

	txs := []*Transaction{coinbase}
	txs = append(txs, bc.Mempool...)

	prevBlock := bc.Blocks[len(bc.Blocks)-1]

	newBlock := NewBlock(
		prevBlock.Index+1,
		txs,
		prevBlock.Hash,
		bc.Difficulty,
	)

	bc.Blocks = append(bc.Blocks, newBlock)
	bc.Mempool = []*Transaction{}

	fmt.Println("New block added to Whytegod (WGD) chain")
}

func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		current := bc.Blocks[i]
		prev := bc.Blocks[i-1]

		if string(current.PrevHash) != string(prev.Hash) {
			return false
		}

		if string(current.Hash) != string(current.HashBlock()) {
			return false
		}
	}
	return true
}