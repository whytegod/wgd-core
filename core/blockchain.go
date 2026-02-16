package core

type Blockchain struct {
	Blocks []Block
}

func NewBlockchain() *Blockchain {
	genesisBlock := NewBlock(0, "Genesis Block", "")
	return &Blockchain{
		Blocks: []Block{genesisBlock},
	}
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(prevBlock.Index+1, data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}