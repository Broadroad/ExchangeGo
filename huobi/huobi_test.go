package huobi

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/ExchangeGo/common"
)

func TestGetTickerWithWs(t *testing.T) {
	hb := NewHuobi(&http.Client{}, "", "", "")
	hb.createWsConn()
	hb.GetTickerWithWs(common.BTC_USDT, func(ticker *common.Ticker) {
		log.Println(ticker)
	})
	time.Sleep(6 * time.Second)
	defer hb.ws.CloseWs()
}
