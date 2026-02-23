package ledger

import "wgd-core/internal/tx"

type UndoEntry struct {
	Key   UTXOKey
	Value tx.TxOutput
}

type BlockUndo struct {
	SpentOutputs []UndoEntry
}