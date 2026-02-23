package consensus

import (
	"errors"
	"math/big"
	"wgd-core/internal/block"
)

func ValidateHeader(h *block.BlockHeader) error {

	hash := h.Hash()
	hashInt := new(big.Int).SetBytes(hash[:])

	target := block.CompactToTarget(h.Bits)

	if hashInt.Cmp(target) > 0 {
		return errors.New("invalid proof of work")
	}

	return nil
}