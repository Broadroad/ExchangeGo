package common

import (
	"time"
)

// Ticker is all exchange share when get ticker return
type Ticker struct {
	ContractType string       `json:"omitempty"`
	Pair         CurrencyPair `json:"pair"`

	// Fcoin
	LastDeal           float64 `json:"lastdeal"`           // last deal price
	LastDealAmount     float64 `json:"lastdealamount"`     // last deal amount
	HighBuy            float64 `json:"highbuy"`            // highest buy one price
	HighBuyAmount      float64 `json:"highbuyamount"`      // highest buy one amount
	LowSell            float64 `json:"lowsell"`            // lowest sell one price
	LowSellAmount      float64 `json:"lowsellamount"`      // lowest buy one amount
	Last24Price        float64 `json:"last24price"`        // price before 24 hours ago
	Last24HighPrice    float64 `json:"last24highprice"`    // highest price in last 24 hours
	Last24LowPrice     float64 `json:"last24lowprice"`     // highest price in last 24 hours
	Last24BaseAmount   float64 `json:"last24baseamount"`   // base coin amount in 24 last hours. btcusdt -> btc
	Last24TargetAmount float64 `json:"last24targetamount"` // target coin amount in last 24 hours. btcusdt -> usdt

	Last24Amount float64 `json:"last24amount"` // last 24 hours volume
	Buy          float64 `json:"buy"`
	Sell         float64 `json:"sell"`
	High         float64 `json:"high"`
	Low          float64 `json:"low"`
	Vol          float64 `json:"vol"`
	Date         uint64  `json:"date"` // unit second
}

type DepthRecord struct {
	Price,
	Amount float64
}

type DepthRecords []DepthRecord

func (dr DepthRecords) Len() int {
	return len(dr)
}

func (dr DepthRecords) Swap(i, j int) {
	dr[i], dr[j] = dr[j], dr[i]
}

func (dr DepthRecords) Less(i, j int) bool {
	return dr[i].Price < dr[j].Price
}

// Depth of CurrencyPair
type Depth struct {
	ContractType string // for future
	Pair         CurrencyPair
	UTime        time.Time
	AskList,
	BidList DepthRecords
}
