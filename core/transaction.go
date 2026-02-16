package core

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"math/big"
)

type TXInput struct {
	TxID      []byte
	OutIndex  int
	Signature []byte
	PubKey    []byte
}

type TXOutput struct {
	Value      uint64
	PubKeyHash []byte
}

type Transaction struct {
	ID      []byte
	Inputs  []TXInput
	Outputs []TXOutput
}

func (tx *Transaction) Hash() []byte {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	_ = enc.Encode(tx)

	hash = sha256.Sum256(encoded.Bytes())
	return hash[:]
}

func (tx *Transaction) SetID() {
	tx.ID = tx.Hash()
}

func NewCoinbaseTx(to string, reward uint64) *Transaction {
	txin := TXInput{TxID: []byte{}, OutIndex: -1}
	txout := TXOutput{Value: reward, PubKeyHash: []byte(to)}

	tx := Transaction{
		Inputs:  []TXInput{txin},
		Outputs: []TXOutput{txout},
	}
	tx.SetID()
	return &tx
}