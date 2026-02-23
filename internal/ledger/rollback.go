package ledger

import "wgd-core/internal/tx"

func (u *UTXOSet) RollbackBlock(txs []*tx.Transaction, undo *BlockUndo) {

	// Remove newly created outputs
	for _, transaction := range txs {
		txHash := transaction.Hash()

		for i := range transaction.Outputs {
			key := UTXOKey{
				TxHash: txHash,
				Index:  uint32(i),
			}
			u.Delete(key)
		}
	}

	// Restore spent outputs
	for _, entry := range undo.SpentOutputs {
		u.Put(entry.Key, entry.Value)
	}
}