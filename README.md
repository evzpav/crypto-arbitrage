# Crypto Arbitrage

It is a command line program to check prices of a pair crypto/BTC (e.g.: NXS/BTC) to the price of same pair in another exchange.
The target amount represents the amount to be arbitraged. The orderbook depth will be checked for that specific amount.

```bash
# Copy configs_example.yaml:
cp configs_example.yaml configs.yaml

# Build binary
make build

# Arguments 
# [coins] [exchanges] [target_btc_amount]
./arbitrage NXS,EOS binance.com,bitfinex.com 0.1

# It will print the spread compared to other exchanges
```
