package main

import (
	"fmt"

	"github.com/evzpav/crypto-arbitrage/pkg/config"
	"github.com/evzpav/crypto-arbitrage/pkg/exchange"
	"github.com/evzpav/crypto-arbitrage/pkg/orderbook"
	goex "github.com/nntaoli-project/GoEx"
)

var apiUser, apiKey string

var exchangeNames = []string{
	goex.BINANCE,
	// goex.CRYPTOPIA,
	goex.BITTREX,
	// goex.BITFINEX,
	goex.HITBTC,
}

var shitcoins = []string{
	"NXS",
	"EOS",
	// "VRC",
	// "BOXX",
	// "GNT",
	// "WAX",
	// "EDO",
	// "NEBL",
}

func main() {
	config, err := config.NewConfig("./configs.yaml")
	if err != nil {
		fmt.Println(err)
	}

	pairs := orderbook.AssembleCurrencyPairs(shitcoins)
	exchanges := exchange.GetExchangeWrappers(config, exchangeNames)
	var arbitrageQuoteTarget = 0.15

	for _, pair := range pairs {
		orderbook.CalculateSpread(arbitrageQuoteTarget, pair, exchanges)
	}

}
