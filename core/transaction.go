package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Transaction struct {
	From      string
	To        string
	Amount    float64
	Signature []byte
	PubKey    []byte
}

func (tx *Transaction) Hash() string {
	record := tx.From + tx.To + fmt.Sprintf("%f", tx.Amount)
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

func (tx *Transaction) Sign(signature []byte, pubKey []byte) {
	tx.Signature = signature
	tx.PubKey = pubKey
}

func (tx *Transaction) Verify() bool {
	return len(tx.Signature) > 0 && len(tx.PubKey) > 0
}

func (tx *Transaction) Print() {
	fmt.Println("From:", tx.From)
	fmt.Println("To:", tx.To)
	fmt.Println("Amount:", tx.Amount)
}