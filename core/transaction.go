package core

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

type Transaction struct {
	From      string
	To        string
	Amount    float64
	Signature string
	PubKey    string
}

func (tx *Transaction) Hash() []byte {
	record := fmt.Sprintf("%s%s%f", tx.From, tx.To, tx.Amount)
	hash := sha256.Sum256([]byte(record))
	return hash[:]
}

func (tx *Transaction) Sign(priv *ecdsa.PrivateKey) {
	hash := tx.Hash()

	r, s, _ := ecdsa.Sign(rand.Reader, priv, hash)
	signature := append(r.Bytes(), s.Bytes()...)

	tx.Signature = hex.EncodeToString(signature)

	pubKey := append(priv.PublicKey.X.Bytes(), priv.PublicKey.Y.Bytes()...)
	tx.PubKey = hex.EncodeToString(pubKey)
}

func (tx *Transaction) Verify() bool {
	if tx.From == "SYSTEM" {
		return true
	}

	sigBytes, _ := hex.DecodeString(tx.Signature)
	pubBytes, _ := hex.DecodeString(tx.PubKey)

	if len(sigBytes) == 0 || len(pubBytes) == 0 {
		return false
	}

	r := big.Int{}
	s := big.Int{}

	sigLen := len(sigBytes)
	r.SetBytes(sigBytes[:sigLen/2])
	s.SetBytes(sigBytes[sigLen/2:])

	x := big.Int{}
	y := big.Int{}

	keyLen := len(pubBytes)
	x.SetBytes(pubBytes[:keyLen/2])
	y.SetBytes(pubBytes[keyLen/2:])

	rawPubKey := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     &x,
		Y:     &y,
	}

	return ecdsa.Verify(&rawPubKey, tx.Hash(), &r, &s)
}