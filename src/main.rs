mod block;
mod blockchain;
mod mining;

use blockchain::Blockchain;
use block::Block;

fn main() {
    println!("WGD Core Node Starting...");

    let mut blockchain = Blockchain::new();

    println!("Mining block 1...");
    let block1 = Block::new(
        1,
        "First WGD block".to_string(),
        blockchain.latest_hash(),
    );

    let mined_block = mining::mine_block(block1, 4);

    blockchain.add_block(mined_block);

    println!("Blockchain length: {}", blockchain.chain.len());
}