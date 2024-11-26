use crate::pb::sf::solana::r#type::v1::AccountBlock;
use crate::pb::sf::substreams::solana::r#type::v1::FilteredAccounts;
use std::collections::HashMap;
use substreams::pb::sf::substreams::index::v1::Keys;
use std::str;
use std::convert::TryInto;

#[substreams::handlers::map]
fn index_accounts(account_block: AccountBlock) -> Result<Keys, substreams::errors::Error> {
    let mut keys = Keys::default();
    let mut account_map = HashMap::new();
    let mut owner_map = HashMap::new();
    let mut buf = [0u8; 44];
    account_block.accounts.iter().for_each(|account| {
        if !account_map.contains_key(&account.address) {
            let len = five8::encode_32(to_32_array(account.address.as_slice()), &mut buf);
            let key = format!("account:{}", str::from_utf8(&buf[..len as usize]).unwrap());
            account_map.insert(account.address.clone(), ());
            keys.keys.push(key);
        }

        if !owner_map.contains_key(&account.owner) {
            let len = five8::encode_32(to_32_array(account.owner.as_slice()), &mut buf);
            let key = format!("owner:{}", str::from_utf8(&buf[..len as usize]).unwrap());
            owner_map.insert(account.owner.clone(), ());
            keys.keys.push(key);
        }
    });

    Ok(keys)
}

#[substreams::handlers::map]
fn filtered_accounts(
    query: String,
    mut account_block: AccountBlock,
) -> Result<FilteredAccounts, substreams::errors::Error> {
    let mut accounts = account_block.accounts;
    let expr_matcher = substreams::expr_matcher(&query);

    let mut buf = [0u8; 44];
     accounts
        .retain(|account| {
            let len = five8::encode_32(to_32_array(account.address.as_slice()), &mut buf);
            let acc_key = format!("account:{}", str::from_utf8(&buf[..len as usize]).unwrap());
            let len = five8::encode_32(to_32_array(account.owner.as_slice()), &mut buf);
            let owner_key = format!("owner:{}", str::from_utf8(&buf[..len as usize]).unwrap());
            return expr_matcher.matches_keys(&[acc_key, owner_key]);
        });

    Ok(FilteredAccounts{accounts})
}

fn to_32_array(slice: &[u8]) -> &[u8; 32] {
    slice.try_into().expect("slice with incorrect length")
}