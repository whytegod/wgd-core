package ledger

import (
	"errors"
	"wgd-core/internal/tx"
)

func (u *UTXOSet) ApplyBlock(txs []*tx.Transaction) (*BlockUndo, error) {

	undo := &BlockUndo{}

	for _, tx := range txs {

		// Spend inputs
		for _, in := range tx.Inputs {

			key := in.PrevTxHash
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
		for i, out := range tx.Outputs {

			newKey := tx.OutputKey(i)
			u.Put(newKey, out.Serialize())
		}
	}

	return undo, nil
}