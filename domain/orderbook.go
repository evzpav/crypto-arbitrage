package domain

import goex "github.com/nntaoli-project/GoEx"

type FillOrderbookReturn struct {
	AveragePrice    float64
	QuoteAmount     float64
	BaseAmount      float64
	PartiallyFilled bool
	Message         string
	Exchange        goex.API
}
type FillOrderbookReturns []FillOrderbookReturn

func (f FillOrderbookReturns) Len() int           { return len(f) }
func (f FillOrderbookReturns) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f FillOrderbookReturns) Less(i, j int) bool { return f[i].AveragePrice < f[j].AveragePrice }
