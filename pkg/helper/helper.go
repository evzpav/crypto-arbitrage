package helper

import (
	"fmt"
	"sort"

	"github.com/evzpav/crypto-arbitrage/domain"
	goex "github.com/nntaoli-project/GoEx"
)

func CalculateSpread(arbitrageQuoteTarget float64, pair goex.CurrencyPair, wrappers []goex.API) {
	var depthSize = 50
	var buyFills, sellFills []domain.FillOrderbookReturn

	for _, ex := range wrappers {
		ob, err := ex.GetDepth(depthSize, pair)
		if err != nil {
			fmt.Printf("error %s: %v\n", ex.GetExchangeName(), err)
			continue
		}
		buyFill := BuyByQuoteCurrency(ob.AskList, arbitrageQuoteTarget)
		sellFill := SellByQuoteCurrency(ob.BidList, arbitrageQuoteTarget)
		buyFill.Exchange = ex
		sellFill.Exchange = ex
		buyFills = append(buyFills, buyFill)
		sellFills = append(sellFills, sellFill)
	}

	sort.Sort(domain.FillOrderbookReturns(buyFills))
	sort.Sort(domain.FillOrderbookReturns(sellFills))

	minBuyFill := buyFills[0]
	maxSellFill := sellFills[len(sellFills)-1]
	fmt.Printf("pair %s \n", pair.ToSymbol("-"))
	fmt.Printf("avgBuys min: %8.8f - %s. %s\n", minBuyFill.AveragePrice, minBuyFill.Exchange.GetExchangeName(), minBuyFill.Message)
	fmt.Printf("avgSells max: %8.8f - %s. %s\n", maxSellFill.AveragePrice, maxSellFill.Exchange.GetExchangeName(), maxSellFill.Message)
	bestSpread := 100 * (minBuyFill.AveragePrice - maxSellFill.AveragePrice) / minBuyFill.AveragePrice
	fmt.Printf("bestSpread: %f \n", bestSpread)

}

func BuyByQuoteCurrency(asks goex.DepthRecords, targetQuoteAmount float64) domain.FillOrderbookReturn {
	return fillOrderbookByQuoteCurrency("Asks quote", asks, targetQuoteAmount)
}

func BuyByBaseCurrency(asks goex.DepthRecords, targetBaseAmount float64) domain.FillOrderbookReturn {
	return fillOrderbookByBaseCurrency("Asks base", asks, targetBaseAmount)
}

func SellByQuoteCurrency(bids goex.DepthRecords, targetQuoteAmount float64) domain.FillOrderbookReturn {
	return fillOrderbookByQuoteCurrency("Bids quote", bids, targetQuoteAmount)
}

func SellByBaseCurrency(bids goex.DepthRecords, targetBaseAmount float64) domain.FillOrderbookReturn {
	return fillOrderbookByBaseCurrency("Bids base", bids, targetBaseAmount)
}

func fillOrderbookByQuoteCurrency(typeDepth string, depthRecords goex.DepthRecords, targetQuoteAmount float64) domain.FillOrderbookReturn {
	var quoteAmount, baseAmount, averagePrice float64
	var f domain.FillOrderbookReturn
	for i, d := range depthRecords {
		quoteAmount += d.Price * d.Amount

		if targetQuoteAmount > quoteAmount {
			baseAmount += d.Amount
			if len(depthRecords) == i+1 { // lastItem on depthRecords array
				f.PartiallyFilled = true
				f.Message = fmt.Sprintf("Not enough %s volume! quoteAmount:%f baseAmount:%f \n", typeDepth, quoteAmount, baseAmount)
				averagePrice = quoteAmount / baseAmount
				break
			}
		} else {
			quoteAmount -= d.Price * d.Amount
			quoteCurrencyDifference := targetQuoteAmount - quoteAmount
			quoteAmount += quoteCurrencyDifference
			baseAmount += quoteCurrencyDifference / d.Price
			break
		}

	}

	averagePrice = quoteAmount / baseAmount
	f.AveragePrice = averagePrice
	f.BaseAmount = baseAmount
	f.QuoteAmount = quoteAmount

	// fmt.Printf("%s - averagePrice: %f, quoteAmount: %f, baseAmount: %f \n", typeDepth, averagePrice, quoteAmount, baseAmount)
	return f
}

func fillOrderbookByBaseCurrency(typeDepth string, depthRecords goex.DepthRecords, targetBaseAmount float64) domain.FillOrderbookReturn {
	var quoteAmount, baseAmount, averagePrice float64
	var f domain.FillOrderbookReturn
	for i, d := range depthRecords {
		baseAmount += d.Amount
		if targetBaseAmount > baseAmount {
			quoteAmount += d.Price * d.Amount
			if len(depthRecords) == i+1 { // lastItem on depthRecords array
				f.PartiallyFilled = true
				f.Message = fmt.Sprintf("Not enough %s volume! quoteAmount:%f baseAmount:%f \n", typeDepth, quoteAmount, baseAmount)
				break
			}
		} else {
			baseAmount -= d.Amount
			baseCurrencyDifference := targetBaseAmount - baseAmount
			baseAmount += baseCurrencyDifference
			quoteAmount += d.Price * baseCurrencyDifference
			break
		}
	}

	averagePrice = quoteAmount / baseAmount
	// fmt.Printf("%s - averagePrice: %f, quoteAmount: %f, baseAmount: %f \n", typeDepth, averagePrice, quoteAmount, baseAmount)
	f.AveragePrice = averagePrice
	f.BaseAmount = baseAmount
	f.QuoteAmount = quoteAmount
	return f
}

func AssembleCurrencyPairs(shitcoins []string) (pairs []goex.CurrencyPair) {
	for _, shitcoin := range shitcoins {
		pairs = append(pairs, goex.NewCurrencyPair2(fmt.Sprintf("%s_BTC", shitcoin)))
	}
	return pairs
}
