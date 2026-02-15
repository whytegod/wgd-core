mod hash;

use std::time::{SystemTime, UNIX_EPOCH};

const MAX_SUPPLY_WGD: u64 = 9_720_000;
const WERTS_PER_WGD: u64 = 100_000_000;
const MAX_SUPPLY: u64 = MAX_SUPPLY_WGD * WERTS_PER_WGD;

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

            if result[0] == 0 && result[1] == 0 {
                break hex::encode(result);
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

        self.mine_reward();
        self.burn(1_000_000); // Burn 0.01 WGD per block
    }

    fn mine_reward(&mut self) {
        let halvings = self.chain.len() / 100;

        let mut reward: u64 = 50 * WERTS_PER_WGD;

        for _ in 0..halvings {
            reward /= 2;
        }

        if reward == 0 {
            println!("⚠️ Mining rewards exhausted.");
            return;
        }

        if self.circulating_supply + reward > MAX_SUPPLY {
            println!("❌ Max supply reached.");
            return;
        }

        self.circulating_supply += reward;

        println!("⛏️ Block Reward: {} Werts", reward);
        println!("💎 Circulating Supply: {}", self.circulating_supply);
        println!("📦 Max Supply: {}", MAX_SUPPLY);
    }

    fn burn(&mut self, amount: u64) {
        if self.circulating_supply >= amount {
            self.circulating_supply -= amount;
            println!("🔥 Burned {} Werts", amount);
            println!("💎 New Circulating Supply: {}", self.circulating_supply);
        }
    }
}

fn main() {
    println!("🚀 WGD CORE — DIGITAL PLATINUM");
    println!("Max Supply: {} WGD", MAX_SUPPLY_WGD);
    println!("1 WGD = {} WERTS", WERTS_PER_WGD);
    println!("----------------------------------");

    let mut blockchain = Blockchain::new();

    for i in 1..=5 {
        blockchain.add_block(format!("Block {}", i));
    }

    println!("----------------------------------");
    println!("Final Circulating Supply: {}", blockchain.circulating_supply);
}