use sha2::{Digest as Sha2Digest, Sha256};
use sha3::{Digest as Sha3Digest, Sha3_256};
use blake3;

fn main() {
    let data = b"WGD - Digital Platinum";

    // ---------------- SHA256 ----------------
    let mut sha256 = Sha256::new();
    sha256.update(data);
    let hash_sha256 = sha256.finalize();

    // ---------------- SHA3-256 ----------------
    let mut sha3 = Sha3_256::new();
    sha3.update(data);
    let hash_sha3 = sha3.finalize();

    // ---------------- BLAKE3 ----------------
    let hash_blake3 = blake3::hash(data);

    println!("----------------------------------------");
    println!("WGD Core Hash Engine");
    println!("----------------------------------------");

    println!("SHA256   : {:x}", hash_sha256);
    println!("SHA3-256 : {:x}", hash_sha3);
    println!("BLAKE3   : {}", hash_blake3.to_hex());

    println!("----------------------------------------");
}