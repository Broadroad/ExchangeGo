package common

import "strings"

var (
	// Currency
	UNKNOWN = Currency{"UNKNOWN", ""}
	USDT    = Currency{"USDT", ""}
	BTC     = Currency{"BTC", ""}
	ETH     = Currency{"ETH", ""}
	EOS     = Currency{"EOS", ""}
	BTM     = Currency{"BTM", ""}

	// Currency pair with usdt
	EOS_USDT = CurrencyPair{EOS, USDT}
	BTC_USDT = CurrencyPair{BTC, USDT}
	BTM_USDT = CurrencyPair{BTM, USDT}

	// Currency pair with btc
	EOS_BTC = CurrencyPair{EOS, BTC}
	BTM_BTC = CurrencyPair{BTM, BTC}

	// Currency pair with eth
	EOS_ETH = CurrencyPair{EOS, ETH}
	BTM_ETH = CurrencyPair{BTM, ETH}

	UNKNOWN_PAIR = CurrencyPair{UNKNOWN, UNKNOWN}
)

// Symbols is a map of BTC_USDT -> btcusdt
type Symbols map[CurrencyPair]string

// ExchangeSymbols is a map of huobi -> Symbols
type ExchangeSymbols map[string]Symbols

// GetExSymbols get symbols of a exchange support
func GetExSymbols(exName string) Symbols {
	ret, ok := exSymbols[exName]
	if !ok {
		return nil
	}
	return ret
}

var exSymbols ExchangeSymbols

// RegisterExsymbol register a symbol to exchange
func RegisterExSymbol(exName string, pair CurrencyPair) {
	if exSymbols == nil {
		exSymbols = make(ExchangeSymbols)
	}

	if _, ok := exSymbols[exName]; !ok {
		exSymbols[exName] = make(Symbols)
	}

	exSymbols[exName][pair] = pair.ToSymbol("")
}

// Currency describe a Currency info
type Currency struct {
	Symbol string
	Desc   string
}

// String return currency symbol
func (c Currency) String() string {
	return c.Symbol
}

// CurrencyPair means A->B(A buy B)
type CurrencyPair struct {
	CurrencyA Currency
	CurrencyB Currency
}

// ToSymbol convert to symbol
func (pair CurrencyPair) ToSymbol(joinChar string) string {
	return strings.Join([]string{pair.CurrencyA.Symbol, pair.CurrencyB.Symbol}, joinChar)
}

// String return BTC_USDT like this
func (c CurrencyPair) String() string {
	return c.ToSymbol("_")
}

// NewCurrency create new Currency
func NewCurrency(symbol, desc string) Currency {
	symbol = strings.ToUpper(symbol)
	switch symbol {
	case "USDT":
		return USDT
	case "BTC":
		return BTC
	case "EOS":
		return EOS
	case "BTM":
		return BTM
	case "ETH":
		return ETH
	default:
		return Currency{strings.ToUpper(symbol), desc}
	}
}

// NewCurrencyPair return new currency pair
func NewCurrencyPair(currencyA Currency, currencyB Currency) CurrencyPair {
	return CurrencyPair{currencyA, currencyB}
}
