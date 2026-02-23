package block

import "wgd-core/internal/tx"

type Block struct {
	Header       BlockHeader
	Transactions []*tx.Transaction
	Height       uint64
}