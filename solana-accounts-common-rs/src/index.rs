use crate::pb::sf::solana::r#type::v1::{Account, AccountBlock};
use crate::pb::sf::substreams::solana::r#type::v1::FilteredAccounts;
use std::collections::HashMap;
use substreams::matches_keys_in_parsed_expr;
use substreams::pb::sf::substreams::index::v1::Keys;

#[substreams::handlers::map]
fn index_accounts(account_block: AccountBlock) -> Result<Keys, substreams::errors::Error> {
    let mut keys = Keys::default();
    let mut account_map = HashMap::new();
    let mut owner_map = HashMap::new();
    account_block.accounts.into_iter().for_each(|account| {
        if !account_map.contains_key(&account.address) {
            let key = format!("account:{}", bs58::encode(&account.address).into_string());
            account_map.insert(account.address, ());
            keys.keys.push(key);
        }

        if !owner_map.contains_key(&account.owner) {
            let key = format!("owner:{}", bs58::encode(&account.owner).into_string());
            owner_map.insert(account.owner, ());
            keys.keys.push(key);
        }
    });

    Ok(keys)
}
#[substreams::handlers::map]
fn filtered_accounts(
    query: String,
    account_block: AccountBlock,
) -> Result<FilteredAccounts, substreams::errors::Error> {
    let mut filtered_accounts = FilteredAccounts::default();
    filtered_accounts.accounts = account_block
        .accounts
        .into_iter()
        .filter(|account| {
            let indexes = index_for_account(&account);
            return matches_keys_in_parsed_expr(&indexes, &query).unwrap();
        })
        .collect();

    Ok(filtered_accounts)
}

fn index_for_account(account: &Account) -> Vec<String> {
    return vec![
        format!("account:{}", bs58::encode(&account.address).into_string()),
        format!("owner:{}", bs58::encode(&account.owner).into_string()),
    ];
}
