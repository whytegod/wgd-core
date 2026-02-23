package node

func (n *Node) reorganize(newTip *BlockNode) {

	// Find fork point
	oldTip := n.BestTip

	if oldTip == nil {
		n.BestTip = newTip
		return
	}

	forkPoint := findFork(oldTip, newTip)

	// Rollback old chain
	current := oldTip
	for current != forkPoint {
		// rollback UTXO here (must implement undo log later)
		current = current.Parent
	}

	// Apply new chain
	path := []*BlockNode{}
	current = newTip
	for current != forkPoint {
		path = append(path, current)
		current = current.Parent
	}

	for i := len(path)-1; i >= 0; i-- {
		// apply block transactions to UTXO
	}

	n.BestTip = newTip
}