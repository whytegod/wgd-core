use crate::state::State;

pub const MAX_SUPPLY: u128 = 9_720_000 * 100_000_000; // 9.72M with 8 decimals

pub struct Consensus {
    pub block_height: u64,
    pub current_reward: u128,
}

impl Consensus {
    pub fn new(initial_reward: u128) -> Self {
        Self {
            block_height: 0,
            current_reward: initial_reward,
        }
    }

    pub fn produce_block(&mut self, state: &mut State, miner: &str) -> Result<(), String> {
        // Enforce emission schedule
        state.mint(miner, self.current_reward, MAX_SUPPLY)?;

        self.block_height += 1;

        // Example reward decay every 210,000 blocks
        if self.block_height % 210_000 == 0 {
            self.current_reward /= 2;
        }

        Ok(())
    }
}