# supermassive

**supermassive** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

## Get started

```
starport serve
```

`serve` command installs dependencies, builds, initializes and starts your blockchain in development.

## Upload contracts

```
supermassived tx wasm store antimatter.wasm --from alice --gas 1500000
supermassived tx wasm store cw721_base.wasm --from alice --gas 1500000
```

View them:
```
supermassived q wasm list-code
```

## Instantiate contracts

```
supermassived tx wasm instantiate 2 '{ "name":"alien", "symbol":"ALIEN", "minter":"cosmos135gzs30un5wg32fyaddlyjgehf4ege850pctq8" }' --from alice --label alien --gas 1000000
```

## Configure

Your blockchain in development can be configured with `config.yml`. To learn more see the [reference](https://github.com/tendermint/starport#documentation).

## Launch

To launch your blockchain live on mutliple nodes use `starport network` commands. Learn more about [Starport Network](https://github.com/tendermint/spn).

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Discord](https://discord.gg/W8trcGV)
