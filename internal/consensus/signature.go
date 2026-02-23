package consensus

import (
	"crypto/sha256"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

func VerifySignature(pubKeyBytes, signature, data []byte) bool {

	pubKey, err := secp256k1.ParsePubKey(pubKeyBytes)
	if err != nil {
		return false
	}

	sig, err := secp256k1.ParseDERSignature(signature)
	if err != nil {
		return false
	}

	hash := sha256.Sum256(data)

	return sig.Verify(hash[:], pubKey)
}