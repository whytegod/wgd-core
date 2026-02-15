use ed25519_dalek::{SigningKey, VerifyingKey, Signature, Signer, Verifier};
use rand::rngs::OsRng;
use sha2::{Sha256, Digest};
use hex;

pub struct Wallet {
    signing_key: SigningKey,
    verifying_key: VerifyingKey,
}

impl Wallet {
    pub fn new() -> Self {
        let mut csprng = OsRng;
        let signing_key = SigningKey::generate(&mut csprng);
        let verifying_key = signing_key.verifying_key();

        Wallet {
            signing_key,
            verifying_key,
        }
    }

    pub fn address(&self) -> String {
        let mut hasher = Sha256::new();
        hasher.update(self.verifying_key.to_bytes());
        hex::encode(hasher.finalize())
    }

    pub fn sign(&self, message: &[u8]) -> Signature {
        self.signing_key.sign(message)
    }

    pub fn verify(
        &self,
        message: &[u8],
        signature: &Signature,
    ) -> bool {
        self.verifying_key.verify(message, signature).is_ok()
    }

    pub fn public_key_bytes(&self) -> Vec<u8> {
        self.verifying_key.to_bytes().to_vec()
    }
}