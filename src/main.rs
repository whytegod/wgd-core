mod hash;

use std::time::{SystemTime, UNIX_EPOCH};

const MAX_SUPPLY_WGD: u64 = 9_720_000;
const WERTS_PER_WGD: u64 = 100_000_000;
const MAX_SUPPLY: u64 = MAX_SUPPLY_WGD * WERTS_PER_WGD;

const INITIAL_REWARD_WGD: u64 = 50;
const HALVING_INTERVAL: u64 = 100;
const BURN_PER_BLOCK: u64 = 1_000_000; // 0.01 WGD

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
            let input = format!("{}{}{}{}{}", index, timestamp, data, previous_hash, nonce);
            let result = hash::wgd_x_hash(input.as_bytes(), nonce);

            // Simple difficulty rule (first two bytes zero)
            if result[0] == 0 && result[1] == 0 {
                break format!("{:x?}", result);
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

        Self {
            chain: vec![genesis],
            circulating_supply: 0,
        }
    }

    fn add_block(&mut self, data: String) {
        let prev = self.chain.last().unwrap();

        let block = Block::new(
            prev.index + 1,
            data,
            prev.hash.clone(),
        );

        self.chain.push(block);

        self.apply_mining_reward();
        self.apply_burn();
    }

    fn apply_mining_reward(&mut self) {
        let halvings = self.chain.len() as u64 / HALVING_INTERVAL;

        let mut reward = INITIAL_REWARD_WGD * WERTS_PER_WGD;

        for _ in 0..halvings {
            reward /= 2;
        }

        if reward == 0 {
            println!("⚠ Mining rewards exhausted.");
            return;
        }

        if self.circulating_supply + reward > MAX_SUPPLY {
            println!("❌ Max supply reached.");
            return;
        }

        self.circulating_supply += reward;

        println!("⛏ Block mined");
        println!("Reward: {} Werts", reward);
        println!("Circulating Supply: {}", self.circulating_supply);
    }

    fn apply_burn(&mut self) {
        if self.circulating_supply >= BURN_PER_BLOCK {
            self.circulating_supply -= BURN_PER_BLOCK;
            println!("🔥 Burned: {} Werts", BURN_PER_BLOCK);
            println!("New Circulating Supply: {}", self.circulating_supply);
        }
    }
}

fn main() {
    println!("WGD CORE — Digital Platinum");
    println!("Max Supply: {} WGD", MAX_SUPPLY_WGD);
    println!("1 WGD = {} WERTS", WERTS_PER_WGD);
    println!("-----------------------------------");

    let mut blockchain = Blockchain::new();

    for i in 1..=5 {
        blockchain.add_block(format!("Block {}", i));
        println!("-----------------------------------");
    }

    println!("Final Circulating Supply: {}", blockchain.circulating_supply);
}