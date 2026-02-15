use crate::block::Block;

pub struct Blockchain {
    pub chain: Vec<Block>,
}

impl Blockchain {
    pub fn new() -> Self {
        let genesis_block = Block::new(0, vec![], String::from("0"));
        Blockchain {
            chain: vec![genesis_block],
        }
    }

    pub fn add_block(&mut self, block: Block) {
        self.chain.push(block);
    }
}