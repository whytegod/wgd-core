package ledger

func (u *UTXOSet) RollbackBlock(txs []*tx.Transaction, undo *BlockUndo) {

	// Remove created outputs
	for _, tx := range txs {
		for i := range tx.Outputs {
			key := tx.OutputKey(i)
			u.Delete(key)
		}
	}

	// Restore spent outputs
	for _, entry := range undo.SpentOutputs {
		u.Put(entry.Key, entry.Value)
	}
}