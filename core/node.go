package core

import (
	"fmt"

	"wgd-core/internal/block"
	"wgd-core/internal/economics"
	"wgd-core/internal/genesis"
	"wgd-core/internal/ledger"
)

type Node struct {
	Chain  []*block.Block
	Ledger *ledger.Ledger
	Supply *economics.Supply
}

func NewNode() *Node {
	l := ledger.NewLedger()
	s := economics.NewSupply(1000000, 50)

	genesisBlock := genesis.CreateGenesisBlock()

	return &Node{
		Chain:  []*block.Block{genesisBlock},
		Ledger: l,
		Supply: s,
	}
}

func (n *Node) AddBlock(transactions []string) {
	prevBlock := n.Chain[len(n.Chain)-1]
	newBlock := block.NewBlock(prevBlock.Index+1, prevBlock.Hash, transactions)

	n.Chain = append(n.Chain, newBlock)

	// Mint reward to miner
	reward := n.Supply.Mint()
	n.Ledger.Credit("miner", reward)
}

func (n *Node) PrintState() {
	fmt.Println("Blockchain length:", len(n.Chain))
	fmt.Println("Total Supply:", n.Supply.TotalSupply)
	fmt.Println("Miner Balance:", n.Ledger.GetBalance("miner"))
}