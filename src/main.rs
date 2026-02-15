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
    nonce: u64,
}

struct Blockchain {
    chain: Vec<Block>,
    circulating_supply: u64,
}

impl Blockchain {
    fn new() -> Self {
        let genesis = Block {
            index: 0,
            timestamp: current_time(),
            nonce: 0,
        };

        Self {
            chain: vec![genesis],
            circulating_supply: 0,
        }
    }

    fn add_block(&mut self) {
        let index = self.chain.len() as u64;

        let block = Block {
            index,
            timestamp: current_time(),
            nonce: 0,
        };

        self.chain.push(block);

        self.apply_reward();
        self.apply_burn();

        println!("⛏ Block {} mined", index);
        println!("Circulating Supply: {}", self.circulating_supply);
        println!("----------------------------------");
    }

    fn apply_reward(&mut self) {
        let halvings = self.chain.len() as u64 / HALVING_INTERVAL;
        let mut reward = INITIAL_REWARD_WGD * WERTS_PER_WGD;

        for _ in 0..halvings {
            reward /= 2;
        }

        if self.circulating_supply + reward <= MAX_SUPPLY {
            self.circulating_supply += reward;
        } else {
            println!("❌ Max supply reached");
        }
    }

    fn apply_burn(&mut self) {
        if self.circulating_supply >= BURN_PER_BLOCK {
            self.circulating_supply -= BURN_PER_BLOCK;
        }
    }
}

fn current_time() -> u128 {
    SystemTime::now()
        .duration_since(UNIX_EPOCH)
        .unwrap()
        .as_millis()
}

fn main() {
    println!("WGD CORE — Digital Platinum");
    println!("Max Supply: {} WGD", MAX_SUPPLY_WGD);
    println!("1 WGD = {} WERTS", WERTS_PER_WGD);
    println!("==================================");

    let mut blockchain = Blockchain::new();

    for _ in 0..5 {
        blockchain.add_block();
    }

    println!("Final Supply: {}", blockchain.circulating_supply);
}