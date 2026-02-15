use sha2::{Sha256, Digest};
use blake3;
use tiny_keccak::{Hasher, Keccak};

const MEMORY_SIZE: usize = 8 * 1024 * 1024; // 8MB

pub fn wgd_x_hash(input: &[u8], nonce: u64) -> [u8; 32] {

    // Phase 1: SHA256
    let mut sha = Sha256::new();
    sha.update(input);
    let h1 = sha.finalize();

    // Phase 2: Blake3
    let h2 = blake3::hash(&h1);

    // Phase 3: Memory Mixing
    let mut buffer = vec![0u8; MEMORY_SIZE];

    for i in 0..MEMORY_SIZE {
        buffer[i] = h2.as_bytes()[i % 32] ^ ((nonce >> (i % 8)) as u8);
    }

    for _ in 0..4 {
        for i in 1..MEMORY_SIZE {
            buffer[i] ^= buffer[i - 1].rotate_left(1);
        }
    }

    // Phase 4: Keccak
    let mut keccak = Keccak::v256();
    keccak.update(&buffer);
    let mut h3 = [0u8; 32];
    keccak.finalize(&mut h3);

    // Phase 5: Final SHA256
    let mut sha_final = Sha256::new();
    sha_final.update(h3);
    let result = sha_final.finalize();

    let mut final_hash = [0u8; 32];
    final_hash.copy_from_slice(&result);

    final_hash
}