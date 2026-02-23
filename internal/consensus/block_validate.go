package consensus

import (
	"errors"
	"wgd-core/internal/block"
	"wgd-core/internal/ledger"
)

func ValidateBlock(
	b *block.Block,
	prevBlock *block.Block,
	utxo *ledger.UTXOSet,
) error {

	// 1. Validate header
	if err := ValidateHeader(&b.Header); err != nil {
		return err
	}

	// 2. Check previous hash
	if b.Header.PrevHash != prevBlock.Header.Hash() {
		return errors.New("invalid previous hash")
	}

	// 3. Validate transactions
	var totalFees uint64

	for i, tx := range b.Transactions {

		if i == 0 {
			continue // coinbase handled separately
		}

		if err := ValidateTransaction(tx, utxo); err != nil {
			return err
		}

		// TODO: calculate fees properly
	}

	// 4. Validate coinbase reward
	coinbase := b.Transactions[0]
	expectedReward := BlockReward(b.Height) + totalFees

	var coinbaseOutputTotal uint64
	for _, out := range coinbase.Outputs {
		coinbaseOutputTotal += out.Value
	}

	if coinbaseOutputTotal > expectedReward {
		return errors.New("coinbase reward exceeds allowed amount")
	}

	return nil
}