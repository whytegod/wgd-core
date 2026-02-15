use std::time::{SystemTime, UNIX_EPOCH};

const MAX_SUPPLY_WGD: u64 = 9_720_000;
const WERTS_PER_WGD: u64 = 100_000_000;
const MAX_SUPPLY: u64 = MAX_SUPPLY_WGD * WERTS_PER_WGD;

const INITIAL_REWARD_WGD: u64 = 50;
const HALVING_INTERVAL: u64 = 100;
const BURN_PER_BLOCK: u64 = 1_000_000; // 0.01 WGD

#[derive(Debug)]
struct Block {
    height: u64,
    timestamp: u128,
}

struct Blockchain {
    blocks: Vec<Block>,
    circulating_supply: u64,
}

impl Blockchain {
    fn new() -> Self {
        let genesis = Block {
            height: 0,
            timestamp: current_time(),
        };

        Self {
            blocks: vec![genesis],
            circulating_supply: 0,
        }
    }

    fn mine_block(&mut self) {
        let height = self.blocks.len() as u64;

        let block = Block {
            height,
            timestamp: current_time(),
        };

        self.blocks.push(block);

        let reward = self.calculate_reward(height);

        if self.circulating_supply + reward <= MAX_SUPPLY {
            self.circulating_supply += reward;
        }

        if self.circulating_supply >= BURN_PER_BLOCK {
            self.circulating_supply -= BURN_PER_BLOCK;
        }

        println!("⛏ Block {} mined", height);
        println!("Reward: {} WERTS", reward);
        println!("Circulating Supply: {}", self.circulating_supply);
        println!("----------------------------------");
    }

    fn calculate_reward(&self, height: u64) -> u64 {
        let halvings = height / HALVING_INTERVAL;
        let mut reward = INITIAL_REWARD_WGD * WERTS_PER_WGD;

        for _ in 0..halvings {
            reward /= 2;
        }

        reward
    }
}

fn current_time() -> u128 {
    SystemTime::now()
        .duration_since(UNIX_EPOCH)
        .unwrap()
        .as_millis()
}

fn main() {
    println!("==================================");
    println!("WGD CORE — Digital Platinum");
    println!("Max Supply: {} WGD", MAX_SUPPLY_WGD);
    println!("1 WGD = {} WERTS", WERTS_PER_WGD);
    println!("==================================");

    let mut blockchain = Blockchain::new();

    for _ in 0..5 {
        blockchain.mine_block();
    }

    println!("Final Circulating Supply: {}", blockchain.circulating_supply);
}