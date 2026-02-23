package ledger

type UndoEntry struct {
	Key   [32]byte
	Value []byte
}

type BlockUndo struct {
	SpentOutputs []UndoEntry
}