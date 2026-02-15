use crate::block::Block;

pub fn mine_block(mut block: Block, difficulty: usize) -> Block {
    let target = "0".repeat(difficulty);

    while &block.hash[..difficulty] != target {
        block.nonce += 1;
        block.hash = block.calculate_hash();
    }

    println!("Block mined: {}", block.hash);
    block
}