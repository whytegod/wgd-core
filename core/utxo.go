package core

import (
	"errors"
)

type UTXOSet struct {
	UTXOs map[string][]TXOutput // txID -> outputs
}

func NewUTXOSet() *UTXOSet {
	return &UTXOSet{
		UTXOs: make(map[string][]TXOutput),
	}
}

// Find spendable outputs for an address
func (u *UTXOSet) FindSpendableOutputs(address string, amount uint64) (uint64, map[string][]int) {
	accumulated := uint64(0)
	spendable := make(map[string][]int)

	for txID, outputs := range u.UTXOs {
		for idx, out := range outputs {
			if out.Address == address && accumulated < amount {
				accumulated += out.Value
				spendable[txID] = append(spendable[txID], idx)

				if accumulated >= amount {
					return accumulated, spendable
				}
			}
		}
	}

	return accumulated, spendable
}

// Remove spent outputs
func (u *UTXOSet) RemoveSpentOutputs(txID string, indexes []int) {
	outputs := u.UTXOs[txID]
	newOutputs := []TXOutput{}

	for i, out := range outputs {
		keep := true
		for _, spentIdx := range indexes {
			if i == spentIdx {
				keep = false
				break
			}
		}
		if keep {
			newOutputs = append(newOutputs, out)
		}
	}

	if len(newOutputs) == 0 {
		delete(u.UTXOs, txID)
	} else {
		u.UTXOs[txID] = newOutputs
	}
}

// Update UTXO set from confirmed block
func (u *UTXOSet) UpdateFromBlock(block *Block) {
	for _, tx := range block.Transactions {

		// Remove spent inputs
		for _, in := range tx.Inputs {
			u.RemoveSpentOutputs(in.TxID, []int{in.OutIndex})
		}

		// Add new outputs
		u.UTXOs[tx.ID] = tx.Outputs
	}
}

// Validate transaction
func (u *UTXOSet) ValidateTransaction(tx *Transaction) error {

	inputTotal := uint64(0)
	outputTotal := uint64(0)

	for _, in := range tx.Inputs {
		outs, exists := u.UTXOs[in.TxID]
		if !exists {
			return errors.New("referenced transaction not found")
		}
		if in.OutIndex >= len(outs) {
			return errors.New("invalid output index")
		}

		inputTotal += outs[in.OutIndex].Value
	}

	for _, out := range tx.Outputs {
		outputTotal += out.Value
	}

	if inputTotal < outputTotal {
		return errors.New("insufficient input value")
	}

	return nil
}