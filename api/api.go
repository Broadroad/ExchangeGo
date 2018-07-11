package api

import "github.com/ExchangeGo/common"
// api interface

type API struct {
	GetTicker(currency common.CurrencyPair) (*Ticker, error)
	GetDepth(size int, currency common.CurrencyPair) (*Depth, error)
	GetKlineRecords(currency common.CurrencyPair, period , size, since int) ([]Kline, error)
	GetTrades(currencyPair common.CurrencyPair, since int64) ([]Trade, error)
	GetExchangeName() string
}