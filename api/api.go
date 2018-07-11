package api

// api interface

type API struct {
	LimitBuy(amount float64, price string, currecny CurrencyPair) (*Order, error)
}