package main

import (
	"fmt"
	"github.com/whytegod/wgd-core/core"
)

func main() {

	bc := core.NewBlockchain()

	miner := core.NewWallet()
	user := core.NewWallet()

	// Mine initial block for miner reward
	bc.MineBlock(miner.Address, []*core.Transaction{})

	// Miner sends coins to user
	tx := core.NewTransaction(miner, user.Address, 10)

	bc.MineBlock(miner.Address, []*core.Transaction{tx})

	fmt.Println("Miner Balance:", bc.Balances[miner.Address])
	fmt.Println("User Balance:", bc.Balances[user.Address])
	fmt.Println("Total Supply:", bc.CurrentSupply)
}