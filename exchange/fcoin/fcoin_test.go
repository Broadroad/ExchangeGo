package fcoin

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/ExchangeGo/common"
)

func TestNewFCoin(t *testing.T) {
	fc := NewFCoin(&http.Client{}, "", "")
	fc.createWsConn()
	fc.GetTickerWithWs(common.BTC_USDT, func(ticker *common.Ticker) {
		log.Println(ticker)
	})
	time.Sleep(6 * time.Second)
	defer fc.ws.CloseWs()
	log.Println(fc.GetExchangeName(), fc.GetTimeOffset())
}
