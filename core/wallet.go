package core

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
	Address    string
}

func NewWallet() *Wallet {
	private, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	hash := sha256.Sum256(pubKey)
	address := hex.EncodeToString(hash[:])

	return &Wallet{
		PrivateKey: private,
		PublicKey:  pubKey,
		Address:    address,
	}
}

func Sign(data []byte, priv *ecdsa.PrivateKey) ([]byte, []byte) {
	hash := sha256.Sum256(data)
	r, s, _ := ecdsa.Sign(rand.Reader, priv, hash[:])
	return r.Bytes(), s.Bytes()
}

func Verify(data, pubKey, rBytes, sBytes []byte) bool {
	hash := sha256.Sum256(data)

	x := big.Int{}
	y := big.Int{}
	keyLen := len(pubKey)
	x.SetBytes(pubKey[:keyLen/2])
	y.SetBytes(pubKey[keyLen/2:])

	rawPubKey := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     &x,
		Y:     &y,
	}

	r := big.Int{}
	s := big.Int{}
	r.SetBytes(rBytes)
	s.SetBytes(sBytes)

	return ecdsa.Verify(&rawPubKey, hash[:], &r, &s)
}