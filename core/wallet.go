package core

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  string
}

func NewWallet() *Wallet {
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	pub := append(priv.PublicKey.X.Bytes(), priv.PublicKey.Y.Bytes()...)

	return &Wallet{
		PrivateKey: priv,
		PublicKey:  hex.EncodeToString(pub),
	}
}