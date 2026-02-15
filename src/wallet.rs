use rand::Rng;

pub struct Wallet {
    pub address: String,
}

impl Wallet {
    pub fn new() -> Self {
        let random: u64 = rand::thread_rng().gen();
        Wallet {
            address: format!("WGD{}", random),
        }
    }
}