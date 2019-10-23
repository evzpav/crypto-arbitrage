package exchange

import (
	"net/http"

	"github.com/evzpav/crypto-arbitrage/pkg/config"
	goex "github.com/nntaoli-project/GoEx"
	"github.com/nntaoli-project/GoEx/binance"
	"github.com/nntaoli-project/GoEx/bitfinex"
	"github.com/nntaoli-project/GoEx/bittrex"
	"github.com/nntaoli-project/GoEx/cryptopia"
	"github.com/nntaoli-project/GoEx/hitbtc"
)

//InitExchange returns pointer to exchange client
func InitExchange(config *config.Config, name string) goex.API {
	pubkey, seckey := config.GetKeys(name)
	// fmt.Printf("init exchange:%s\n", name)
	// fmt.Printf("pubkey:%s, seckey: %s\n", pubkey, seckey)

	switch name {
	case goex.BINANCE:
		return binance.New(http.DefaultClient, pubkey, seckey)
	case goex.BITFINEX:
		return bitfinex.New(http.DefaultClient, pubkey, seckey)
	// case BITMEX:
	// 	return bitmex.New(http.DefaultClient, pubkey, seckey)
	case goex.BITTREX:
		return bittrex.New(http.DefaultClient, pubkey, seckey)
	case goex.CRYPTOPIA:
		return cryptopia.New(http.DefaultClient, pubkey, seckey)
	case goex.HITBTC:
		return hitbtc.New(http.DefaultClient, pubkey, seckey)
	default:
		return nil
	}
}

//GetExchangeWrappers get wrappers based on slice of names
func GetExchangeWrappers(config *config.Config, exchanges []string) []goex.API {
	wrappers := make([]goex.API, len(exchanges))
	for i, ex := range exchanges {
		wrappers[i] = InitExchange(config, ex)
	}
	return wrappers
}
