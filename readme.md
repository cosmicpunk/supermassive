# supermassive

**supermassive** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

## Get started

```sh
starport serve
```

`serve` command installs dependencies, builds, initializes and starts your blockchain in development.

## Upload contracts

```sh
supermassived tx wasm store antimatter.wasm --from alice --gas 1500000
supermassived tx wasm store cw721_base.wasm --from alice --gas 1500000
```

View them:
```sh
supermassived q wasm list-code
```

## Instantiate contracts

```sh
supermassived tx wasm instantiate 1 '{ "name":"supermassive" }' --from alice --label marketplace --gas 1000000
supermassived tx wasm instantiate 2 '{ "name":"alien", "symbol":"ALIEN", "minter":"cosmos135gzs30un5wg32fyaddlyjgehf4ege850pctq8" }' --from alice --label alien --gas 1000000
```

## Perform Actions

Mint NFT:
```sh
supermassived tx wasm execute cosmos1qxxlalvsdjd07p07y3rc5fu6ll8k4tmecu7e9y '{
  "mint": {
    "name": "Alien 1",
    "owner": "cosmos135gzs30un5wg32fyaddlyjgehf4ege850pctq8",
    "token_id": "1"
  }
}'  -y --from alice --gas 1000000
```

List NFT in marketplace:
```sh
supermassived tx wasm execute cosmos1qxxlalvsdjd07p07y3rc5fu6ll8k4tmecu7e9y '{
  "send_nft": {
    "contract": "cosmos1hqrdl6wstt8qzshwc6mrumpjk9338k0lr4dqxd",
    "token_id": "1",
    "msg": "eyAibGlzdF9wcmljZSI6IHsgImFtb3VudCI6ICIxMDAwMDAiLCAiZGVub20iOiAic3Rha2UiIH19"
  }
}' -y --from alice --gas 1000000
```
This calls `send_nft` in the cw721 contract which passes along a BASE64 encoded msg to the marketplace contract.
The encoded msg is: `{ "list_price": { "amount": "100000", "denom": "stake" }}`. The marketplace contract receives
the message and offers it up for sale at the list price.


## Configure

Your blockchain in development can be configured with `config.yml`. To learn more see the [reference](https://github.com/tendermint/starport#documentation).

## Launch

To launch your blockchain live on mutliple nodes use `starport network` commands. Learn more about [Starport Network](https://github.com/tendermint/spn).

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Discord](https://discord.gg/W8trcGV)
