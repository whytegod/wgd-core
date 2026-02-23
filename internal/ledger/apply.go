package ledger

import (
	"errors"
	"wgd-core/internal/tx"
)

func (u *UTXOSet) ApplyBlock(txs []*tx.Transaction) (*BlockUndo, error) {

	undo := &BlockUndo{}

	for _, transaction := range txs {

		txHash := transaction.Hash()

		// Spend inputs
		for _, in := range transaction.Inputs {

			key := UTXOKey{
				TxHash: in.PrevTxHash,
				Index:  in.OutputIndex,
			}

			value, exists := u.Get(key)
			if !exists {
				return nil, errors.New("utxo not found")
			}

			undo.SpentOutputs = append(undo.SpentOutputs, UndoEntry{
				Key:   key,
				Value: value,
			})

			u.Delete(key)
		}

		// Add outputs
		for i, out := range transaction.Outputs {

			newKey := UTXOKey{
				TxHash: txHash,
				Index:  uint32(i),
			}

			u.Put(newKey, out)
		}
	}

	return undo, nil
}