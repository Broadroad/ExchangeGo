package fcoin

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/ExchangeGo/errors"

	. "github.com/ExchangeGo/api"
	. "github.com/ExchangeGo/common"
	. "github.com/ExchangeGo/utils"
	. "github.com/ExchangeGo/websocket"
)

const (
	DEPTH_API                 = "market/depth/%s/%s"
	TRADE_URL                 = "orders"
	GET_ACCOUNT_API           = "accounts/balance"
	GET_ORDER_API             = "orders/%s"
	GET_UNFINISHED_ORDERS_API = "getUnfinishedOrdersIgnoreTradeType"
	PLACE_ORDER_API           = "order"
	WITHDRAW_API              = "withdraw"
	CANCELWITHDRAW_API        = "cancelWithdraw"
	SERVER_TIME               = "public/server-time"
)

// FCoinTicker is fcoin return data format
type FCoinTicker struct {
	Ticker
	SellAmount,
	BuyAmount float64
}

// FCoin can get fcoin.com data
type FCoin struct {
	ws         *WsConn
	httpClient *http.Client
	baseUrl,
	accessKey,
	secretKey string
	timeoffset        int64 //server timestamp
	wsTickerHandleMap map[string]func(*Ticker)
}

// NewFCoin new a fcoin client
func NewFCoin(client *http.Client, apikey, secretkey string) *FCoin {
	fc := &FCoin{baseUrl: "https://api.fcoin.com/v2/", accessKey: apikey, secretKey: secretkey, httpClient: client}
	fc.wsTickerHandleMap = make(map[string]func(*Ticker))
	fc.setTimeOffset()
	return fc
}

// createWsConn create a fcoin websocket
func (fc *FCoin) createWsConn() {
	if fc.ws != nil {
		return
	}

	fc.ws = NewWsConn("wss://api.fcoin.com/v2/ws")

	fc.ws.ReConnect()
	fc.ws.ReceiveMessage(func(msg []byte) {
		datamap := make(map[string]interface{})
		err := json.Unmarshal(msg, &datamap)
		if err != nil {
			log.Println("json unmarshal error for ", string(msg))
			return
		}

		fmt.Println(datamap)
		if datamap["type"] != nil {
			tp, isok := datamap["type"].(string)
			if !isok {
				log.Print(errors.New("error when convert type"))
				return
			}

			switch {
			case tp == "hello":
				ts, isok := datamap["ts"].(float64)
				fmt.Println(ts)
				if !isok {
					log.Print(errors.New("error when convert ts"))
					return
				}
				fc.timeoffset = int64(ts)

			case tp == "topics":
				topics, isok := datamap["topics"].([]string)
				if !isok {
					log.Print(errors.New("error when convert topics"))
					return
				}
				for topic := range topics {
					log.Print("subscribe topic: {}", topic)

				}

			case strings.Contains(tp, "ticker"):
				log.Print("message type is ", tp)
				tickmap := datamap["ticker"].([]interface{})

				ticker := new(FCoinTicker)
				ticker.Pair = CurrencyPairMap[tp]
				ticker.Date = uint64(time.Now().Nanosecond() / 1000)

				ticker.LastDeal = ToFloat64(tickmap[0])
				ticker.LastDealAmount = ToFloat64(tickmap[1])
				ticker.HighBuy = ToFloat64(tickmap[2])
				ticker.HighBuyAmount = ToFloat64(tickmap[3])
				ticker.LowSell = ToFloat64(tickmap[4])
				ticker.LowSellAmount = ToFloat64(tickmap[5])
				ticker.Last24Price = ToFloat64(tickmap[6])
				ticker.Last24HighPrice = ToFloat64(tickmap[7])
				ticker.Last24LowPrice = ToFloat64(tickmap[8])
				ticker.Last24Amount = ToFloat64(tickmap[9])
				ticker.Last24BaseAmount = ToFloat64(tickmap[10])
				fmt.Println(ticker.LastDeal)
			}
		}
	})
}

// GetTickerWithWs get ticker with websocket
func (fc *FCoin) GetTickerWithWs(pair CurrencyPair, handle func(ticker *Ticker)) error {
	fc.createWsConn()
	topic := fmt.Sprintf("ticker.%s", strings.ToLower(pair.ToSymbol("")))
	fmt.Println(topic)
	fc.wsTickerHandleMap[topic] = handle
	channel := []string{"ticker.btcusdt"}
	return fc.ws.Subscribe(map[string]interface{}{
		"cmd":  "sub",
		"args": channel,
		"id":   channel})
}

// setTimeOffset get server timestamp, because server and client time can not exceed 30 seconds
func (ft *FCoin) setTimeOffset() error {
	respmap, err := HttpGet(ft.httpClient, ft.baseUrl+"public/server-time")
	if err != nil {
		return err
	}
	stime := int64(ToInt(respmap["data"]))
	st := time.Unix(stime/1000, 0)
	lt := time.Now()
	offset := st.Sub(lt).Seconds()
	ft.timeoffset = int64(offset)
	return nil
}

// GetTicker get ticker data
func (ft *FCoin) GetTicker(currencyPair CurrencyPair) (*FCoinTicker, error) {
	respmap, err := HttpGet(ft.httpClient, ft.baseUrl+fmt.Sprintf("market/ticker/%s",
		strings.ToLower(currencyPair.ToSymbol(""))))

	if err != nil {
		return nil, err
	}

	////log.Println("ticker respmap:", respmap)
	if respmap["status"].(float64) != 0 {
		return nil, errors.New(respmap["err-msg"].(string))
	}

	//
	tick, ok := respmap["data"].(map[string]interface{})
	if !ok {
		return nil, API_ERR
	}

	tickmap, ok := tick["ticker"].([]interface{})
	if !ok {
		return nil, API_ERR
	}

	ticker := new(FCoinTicker)
	ticker.Pair = currencyPair
	ticker.Date = uint64(time.Now().Nanosecond() / 1000)

	ticker.LastDeal = ToFloat64(tickmap[0])
	ticker.LastDealAmount = ToFloat64(tickmap[1])
	ticker.HighBuy = ToFloat64(tickmap[2])
	ticker.HighBuyAmount = ToFloat64(tickmap[3])
	ticker.LowSell = ToFloat64(tickmap[4])
	ticker.LowSellAmount = ToFloat64(tickmap[5])
	ticker.Last24Price = ToFloat64(tickmap[6])
	ticker.Last24HighPrice = ToFloat64(tickmap[7])
	ticker.Last24LowPrice = ToFloat64(tickmap[8])
	ticker.Last24Amount = ToFloat64(tickmap[9])
	ticker.Last24BaseAmount = ToFloat64(tickmap[10])

	return ticker, nil
}

// GetDepth get the depth of the currency pair
func (ft *FCoin) GetDepth(size int, currency CurrencyPair) (*Depth, error) {
	respmap, err := HttpGet(ft.httpClient, ft.baseUrl+fmt.Sprintf("market/depth/L20/%s", strings.ToLower(currency.ToSymbol(""))))
	if err != nil {
		return nil, err
	}

	if respmap["status"].(float64) != 0 {
		return nil, errors.New(respmap["err-msg"].(string))
	}

	datamap := respmap["data"].(map[string]interface{})

	bids, ok1 := datamap["bids"].([]interface{})
	asks, ok2 := datamap["asks"].([]interface{})

	if !ok1 || !ok2 {
		return nil, errors.New("depth error")
	}

	depth := new(Depth)
	depth.Pair = currency

	n := 0
	for i := 0; i < len(bids); {
		depth.BidList = append(depth.BidList, DepthRecord{ToFloat64(bids[i]), ToFloat64(bids[i+1])})
		i += 2
		n++
		if n == size {
			break
		}
	}

	n = 0
	for i := 0; i < len(asks); {
		depth.AskList = append(depth.AskList, DepthRecord{ToFloat64(asks[i]), ToFloat64(asks[i+1])})
		i += 2
		n++
		if n == size {
			break
		}
	}

	//sort.Sort(sort.Reverse(depth.AskList))
	return depth, nil
}

// GetExchangeName return ExchangeName
func (fc *FCoin) GetExchangeName() string {
	return FCOIN
}

// GetTimeOffset
func (fc *FCoin) GetTimeOffset() int64 {
	return fc.timeoffset
}
