package wallet

import (
	"crypto/rand"
	"crypto/sha256"

	"golang.org/x/crypto/ripemd160"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

type KeyPair struct {
	PrivateKey *secp256k1.PrivateKey
	PublicKey  []byte
}

func NewKeyPair() (*KeyPair, error) {
	privKey, err := secp256k1.GeneratePrivateKey()
	if err != nil {
		return nil, err
	}

	pubKey := privKey.PubKey().SerializeCompressed()

	return &KeyPair{
		PrivateKey: privKey,
		PublicKey:  pubKey,
	}, nil
}

func HashPubKey(pubKey []byte) []byte {
	sha := sha256.Sum256(pubKey)

	ripemd := ripemd160.New()
	ripemd.Write(sha[:])
	return ripemd.Sum(nil)
}