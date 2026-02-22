package core

import (
	"bytes"
	"encoding/gob"
)

type Transaction struct {
	From      string
	To        string
	Amount    int
	PublicKey []byte
	R         []byte
	S         []byte
}

func (tx *Transaction) Hash() []byte {
	var encoded bytes.Buffer
	gob.NewEncoder(&encoded).Encode(tx.From + tx.To)
	return encoded.Bytes()
}

func NewCoinbaseTx(to string, amount int) *Transaction {
	return &Transaction{
		From:   "COINBASE",
		To:     to,
		Amount: amount,
	}
}

func NewTransaction(fromWallet *Wallet, to string, amount int) *Transaction {
	tx := &Transaction{
		From:      fromWallet.Address,
		To:        to,
		Amount:    amount,
		PublicKey: fromWallet.PublicKey,
	}

	r, s := Sign(tx.Hash(), fromWallet.PrivateKey)
	tx.R = r
	tx.S = s

	return tx
}