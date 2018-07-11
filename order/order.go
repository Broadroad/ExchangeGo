package order

import "github.com/ExchangeGo/common"

const (
	BUY = 1 + iota
	SELL
	BUY_MARKET
	SELL_MARKET
)

type TradeSide int

func (ts TradeSide) String() string {
	switch ts {
	case 1:
		return "BUY"
	case 2:
		return "SELL"
	case 3:
		return "BUY_MARKET"
	case 4:
		return "SELL_MARKET"
	default:
		return "UNKNOWN"
	}
}

type TradeStatus int

var tradeStatusSymbol = [...]string{"UNFINISH", "PART_FINISH", "FINISH", "CANCEL", "REJECT", "CANCEL_ING"}

func (ts TradeStatus) String() string {
	return tradeStatusSymbol[ts]
}

const (
	ORDER_UNFINISH = iota
	ORDER_PART_FINISH
	ORDER_FINISH
	ORDER_CANCEL
	ORDER_REJECT
	ORDER_CANCEL_ING
)

type Order struct {
	Price,
	Amount,
	AvgPrice,
	DealAmount,
	Fee float64
	OrderID2  string
	OrderID   int
	OrderTime int
	Status    TradeStatus
	Currency  common.CurrencyPair
	Side      TradeSide
}
