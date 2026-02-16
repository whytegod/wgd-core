package core

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
)

type Transaction struct {
	From      string
	To        string
	Amount    float64
	Signature string
	Hash      string
}

func (tx *Transaction) CalculateHash() string {
	record := tx.From + tx.To + string(rune(tx.Amount))
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

func (tx *Transaction) SignTransaction(privateKey *ecdsa.PrivateKey) {
	hash := sha256.Sum256([]byte(tx.CalculateHash()))

	r, s, _ := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	signature := append(r.Bytes(), s.Bytes()...)
	tx.Signature = hex.EncodeToString(signature)
}

func GenerateKeyPair() (*ecdsa.PrivateKey, string) {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	pubKey := append(
		privateKey.PublicKey.X.Bytes(),
		privateKey.PublicKey.Y.Bytes()...,
	)

	address := sha256.Sum256(pubKey)
	return privateKey, hex.EncodeToString(address[:])
}