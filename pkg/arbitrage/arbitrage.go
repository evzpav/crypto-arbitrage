package arbitrage

import (
	"fmt"
	"log"

	"github.com/evzpav/crypto-arbitrage/pkg/config"
	"github.com/evzpav/crypto-arbitrage/pkg/orderbook"
)

type Arbitrage struct {
	Config *config.Config
}

func New(configPath string) (*Arbitrage, error) {
	config, err := config.NewConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("Could not load configs: %v", err)
	}
	return &Arbitrage{
		Config: config,
	}, nil
}

func (arb *Arbitrage) Run(exchangesNames, cryptos []string, target float64) {
	pairs := orderbook.AssembleCurrencyPairs(cryptos)
	if len(pairs) == 0 {
		log.Fatal("Failed to assemble pairs")
	}

	exchangeWrappers := arb.Config.GetExchangeWrappers(exchangesNames)
	if len(exchangeWrappers) == 0 {
		log.Fatal("Failed to get exchange wrappers")
	}

	for _, pair := range pairs {
		orderbook.CalculateSpread(target, pair, exchangeWrappers)
	}

}
