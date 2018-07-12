package api

import "github.com/ExchangeGo/common"
// api interface

type APIERROR struct {
	ErrCode
	ErrMsg
	OriginErrMsg string
}

func (e ApiError) Error() string {
	return e.ErrMsg
}

var (
	// API_ERR is returned when call api error happens
	API_ERR 		= ApiError{ErrCode: "EX_ERR_0000", ErrMsg: "unknown error"}
	HTTP_ERR_CODE 	= ApiError{ErrCode: "HTTP_ERR_0001", ErrMsg: "http request error"}
)

//NewSuccessResponse success response
func NewSuccessResponse(data interface{}) Response {
	return Response{Status: SUCCESS, Data: data}
}

type API struct {
	GetTicker(currency common.CurrencyPair) (*Ticker, error)
	GetDepth(size int, currency common.CurrencyPair) (*Depth, error)
	GetKlineRecords(currency common.CurrencyPair, period , size, since int) ([]Kline, error)
	GetTrades(currencyPair common.CurrencyPair, since int64) ([]Trade, error)
	GetExchangeName() string
}