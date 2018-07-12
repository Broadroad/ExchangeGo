package api

import "github.com/ExchangeGo/common"

// api interface

type ApiError struct {
	ErrCode,
	ErrMsg,
	OriginErrMsg string
}

func (e ApiError) Error() string {
	return e.ErrMsg
}

var (
	// API_ERR is returned when call api error happens
	API_ERR       = ApiError{ErrCode: "EX_ERR_0000", ErrMsg: "unknown error"}
	HTTP_ERR_CODE = ApiError{ErrCode: "HTTP_ERR_0001", ErrMsg: "http request error"}
)

type API interface {
	GetTicker(currency common.CurrencyPair) (*common.Ticker, error)
	GetDepth(size int, currency common.CurrencyPair) (*common.Depth, error)
	GetExchangeName() string
}
