use crate::block::Block;

pub struct Blockchain {
    pub chain: Vec<Block>,
}

impl Blockchain {
    pub fn new() -> Self {
        let genesis = Block::new(
            0,
            "Genesis Block".to_string(),
            "0".to_string(),
        );

        Blockchain {
            chain: vec![genesis],
        }
    }

    pub fn add_block(&mut self, block: Block) {
        self.chain.push(block);
    }

    pub fn latest_hash(&self) -> String {
        self.chain.last().unwrap().hash.clone()
    }
}