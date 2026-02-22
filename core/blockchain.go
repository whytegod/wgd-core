package core

import "fmt"

const MaxSupply = 9720000
const HalvingInterval = 1051200
const InitialReward = 50

type Blockchain struct {
	Blocks        []*Block
	CurrentSupply int
	Balances      map[string]int
}

func NewBlockchain() *Blockchain {
	genesis := NewBlock([]*Transaction{}, "0", 0)

	return &Blockchain{
		Blocks:        []*Block{genesis},
		CurrentSupply: 0,
		Balances:      make(map[string]int),
	}
}

func (bc *Blockchain) getReward() int {
	height := len(bc.Blocks)
	halvings := height / HalvingInterval

	reward := InitialReward
	for i := 0; i < halvings; i++ {
		reward /= 2
	}

	if reward <= 0 {
		return 0
	}

	if bc.CurrentSupply+reward > MaxSupply {
		return MaxSupply - bc.CurrentSupply
	}

	return reward
}

func (bc *Blockchain) ValidateTransaction(tx *Transaction) bool {

	if tx.From == "COINBASE" {
		return true
	}

	if bc.Balances[tx.From] < tx.Amount {
		return false
	}

	if !Verify(tx.Hash(), tx.PublicKey, tx.R, tx.S) {
		return false
	}

	return true
}

func (bc *Blockchain) MineBlock(miner string, transactions []*Transaction) {

	validTxs := []*Transaction{}

	for _, tx := range transactions {
		if bc.ValidateTransaction(tx) {
			validTxs = append(validTxs, tx)
		}
	}

	reward := bc.getReward()
	if reward > 0 {
		coinbase := NewCoinbaseTx(miner, reward)
		validTxs = append(validTxs, coinbase)
		bc.CurrentSupply += reward
		bc.Balances[miner] += reward
	}

	prev := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(validTxs, prev.Hash, len(bc.Blocks))

	bc.Blocks = append(bc.Blocks, newBlock)

	for _, tx := range validTxs {
		if tx.From != "COINBASE" {
			bc.Balances[tx.From] -= tx.Amount
		}
		bc.Balances[tx.To] += tx.Amount
	}

	fmt.Printf("Block %d mined | Supply: %d\n",
		newBlock.Index,
		bc.CurrentSupply)
}