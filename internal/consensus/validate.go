package consensus

import (
	"errors"
	"wgd-core/internal/ledger"
	"wgd-core/internal/tx"
)

func ValidateTransaction(t *tx.Transaction, utxo *ledger.UTXOSet) error {

	var totalInput uint64
	var totalOutput uint64

	for _, in := range t.Inputs {

		prevOut, ok := utxo.Get(in.PrevTxHash, in.OutputIndex)
		if !ok {
			return errors.New("invalid input: UTXO not found")
		}

		totalInput += prevOut.Value

		// TODO: verify signature here (ECDSA)
	}

	for _, out := range t.Outputs {
		totalOutput += out.Value
	}

	if totalOutput > totalInput {
		return errors.New("output exceeds input")
	}

	return nil
}