package common

type Ticker struct {
	ContractType string       `json:"omitempty"`
	Pair         CurrencyPair `json:"omitempty"`
	Last24Amount float64      `json:last24amount` // 过去24小时成交量
	LastAmount	 float64 	  `json:lastamount` // 上一笔成交
	Last         float64      `json:"last"` // 最新成交价
	Buy          float64      `json:"buy"`
	Sell         float64      `json:"sell"`
	High         float64      `json:"high"`
	Low          float64      `json:"low"`
	Vol          float64      `json:"vol"`
	Date         uint64       `json:"date"` // unit second
}