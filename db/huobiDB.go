package db

import (
	"time"
)

// each pair and each time period has a table
// Ticker is common ticker table
type Ticker struct {
	ID                 int `gorm:"primary_key"`
	CreatedAt          time.Time
	LastDeal           float64 `gorm:"column:last_deal"`
	LastDealAmount     float64 `gorm "column:last_deal_amount"`
	HighBuy            float64 `gorm "column:high_buy"`
	HighBuyAmount      float64 `gorm: "column:high_buy_amount"`
	LowSell            float64 `gorm: "column:low_sell"`
	LowSellAmount      float64 `gorm: "column:low_sell_amount"`
	Last24Price        float64 `gorm: "column:last_24_price"`
	Last24HighPrice    float64 `gorm: "column:last_24_high_price"`
	Last24LowPrice     float64 `gorm: "column:last_24_low_price"`
	Last24BaseAmount   float64 `gorm: "column:last_24_base_amount"`
	Last24TargetAmount float64 `gorm: "column:last_24_target_amount"`
	timestamp          time.Time
}
