use std::collections::HashMap;

pub struct State {
    pub balances: HashMap<String, u128>,
    pub total_supply: u128,
}

impl State {
    pub fn new() -> Self {
        Self {
            balances: HashMap::new(),
            total_supply: 0,
        }
    }

    pub fn mint(&mut self, address: &str, amount: u128, max_supply: u128) -> Result<(), String> {
        if self.total_supply + amount > max_supply {
            return Err("Max supply exceeded".to_string());
        }

        let balance = self.balances.entry(address.to_string()).or_insert(0);
        *balance += amount;
        self.total_supply += amount;

        Ok(())
    }

    pub fn transfer(&mut self, from: &str, to: &str, amount: u128) -> Result<(), String> {
        let sender_balance = self.balances.entry(from.to_string()).or_insert(0);

        if *sender_balance < amount {
            return Err("Insufficient balance".to_string());
        }

        *sender_balance -= amount;

        let receiver_balance = self.balances.entry(to.to_string()).or_insert(0);
        *receiver_balance += amount;

        Ok(())
    }
}