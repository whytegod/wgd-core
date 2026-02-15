package core

type Blockchain struct {
	Blocks []Block
}

func NewBlockchain() *Blockchain {
	genesis := NewBlock(0, []Transaction{}, "0")

	return &Blockchain{
		Blocks: []Block{genesis},
	}
}

func (bc *Blockchain) AddBlock(txs []Transaction) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(prevBlock.Index+1, txs, prevBlock.Hash)

	bc.Blocks = append(bc.Blocks, newBlock)
}