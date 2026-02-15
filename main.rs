mod state;
mod consensus;

use state::State;
use consensus::{Consensus, MAX_SUPPLY};

fn main() {
    let mut state = State::new();

    // Initial reward: 50 WGD (8 decimals like BTC style)
    let initial_reward: u128 = 50 * 100_000_000;

    let mut consensus = Consensus::new(initial_reward);

    let miner_address = "WGD_MINER_001";

    // Simulate 5 blocks
    for _ in 0..5 {
        match consensus.produce_block(&mut state, miner_address) {
            Ok(_) => {
                println!("Block {} mined successfully", consensus.block_height);
            }
            Err(e) => {
                println!("Block rejected: {}", e);
            }
        }
    }

    println!("Total Supply: {}", state.total_supply);
    println!("Max Supply: {}", MAX_SUPPLY);
}