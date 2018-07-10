package common

import (
	"strings"
	"time"
)

type Ticker struct {
	ContractType string       `json:"omitempty"`
	Pair         CurrencyPair `json:"omitempty"`
	Last24Amount float64      `json:last24amount` // last 24 hours volume
	LastAmount   float64      `json:lastamount`   // last deal volume
	Last         float64      `json:"last"`       // last deal price
	Buy          float64      `json:"buy"`
	Sell         float64      `json:"sell"`
	High         float64      `json:"high"`
	Low          float64      `json:"low"`
	Vol          float64      `json:"vol"`
	Date         uint64       `json:"date"` // unit second
}

type Currency struct {
	Symbol string
	Desc   string
}

func (c Currency) String() string {
	return c.Symbol
}

// A->B(A buy B)
type CurrencyPair struct {
	CurrencyA Currency
	CurrencyB Currency
}

// ToSymbol convert to symbol
func (pair CurrencyPair) ToSymbol(joinChar string) string {
	return strings.Join([]string{pair.CurrencyA.Symbol, pair.CurrencyB.Symbol}, joinChar)
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
