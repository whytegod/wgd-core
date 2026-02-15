mod hash;

use std::time::{SystemTime, UNIX_EPOCH};

const MAX_SUPPLY: u64 = 9_720_000;          // 9,720,000 WGD fixed supply
const WERTS_PER_WGD: u64 = 100_000_000;     // Smallest unit = WERT

#[derive(Clone)]
struct Block {
    index: u64,
    timestamp: u128,
    data: String,
    previous_hash: String,
    nonce: u64,
    hash: String,
}

impl Block {
    fn new(index: u64, data: String, previous_hash: String) -> Self {
        let timestamp = SystemTime::now()
            .duration_since(UNIX_EPOCH)
            .unwrap()
            .as_millis();

        let mut nonce = 0;

        let hash = loop {
            let block_data = format!("{}{}{}{}{}", index, timestamp, data, previous_hash, nonce);
            let result = hash::wgd_x_hash(block_data.as_bytes(), nonce);

            if result.starts_with("0000") {   // Difficulty level
                break result;
            }

            nonce += 1;
        };

        Block {
            index,
            timestamp,
            data,
            previous_hash,
            nonce,
            hash,
        }
    }
}

struct Blockchain {
    chain: Vec<Block>,
    circulating_supply: u64,
}

impl Blockchain {
    fn new() -> Self {
        let genesis = Block::new(0, "Genesis Block - Digital Platinum".into(), "0".into());

        Blockchain {
            chain: vec![genesis],
            circulating_supply: 0,
        }
    }

    fn add_block(&mut self, data: String) {
        let previous_block = self.chain.last().unwrap();
        let new_block = Block::new(
            previous_block.index + 1,
            data,
            previous_block.hash.clone(),
        );

        self.chain.push(new_block);
    }

    fn mint(&mut self, amount_wgd: u64) {
        if self.circulating_supply + amount_wgd > MAX_SUPPLY {
            println!("❌ Cannot mint beyond fixed supply of {} WGD", MAX_SUPPLY);
        } else {
            self.circulating_supply += amount_wgd;
            println!("✅ Minted {} WGD", amount_wgd);
            println!("💎 Circulating Supply: {} / {}", self.circulating_supply, MAX_SUPPLY);
        }
    }

    fn print_chain(&self) {
        for block in &self.chain {
            println!("-------------------------------");
            println!("Index: {}", block.index);
            println!("Timestamp: {}", block.timestamp);
            println!("Data: {}", block.data);
            println!("Nonce: {}", block.nonce);
            println!("Hash: {}", block.hash);
            println!("Previous Hash: {}", block.previous_hash);
        }
    }
}

fn main() {
    println!("🚀 WGD CORE — DIGITAL PLATINUM");
    println!("Fixed Supply: {} WGD", MAX_SUPPLY);
    println!("Smallest Unit: 1 WGD = {} WERTS", WERTS_PER_WGD);
    println!("----------------------------------");

    let mut blockchain = Blockchain::new();

    blockchain.mint(1_000_000);
    blockchain.add_block("First Transaction".into());

    blockchain.mint(500_000);
    blockchain.add_block("Second Transaction".into());

    blockchain.print_chain();
}