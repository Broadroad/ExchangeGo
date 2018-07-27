package huobi

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	. "github.com/ExchangeGo/common"
	. "github.com/ExchangeGo/websocket"
)

// Huobi is the struct to buy/sell/subscribe
type Huobi struct {
	baseUrl           string
	ws                *WsConn
	httpClient        *http.Client
	apikey            string
	secretkey         string
	accountId         string
	wsTickerHandleMap map[string]func(*Ticker)
}

//NewHuobi return a new Huobi
func NewHuobi(client *http.Client, apikey, secretkey, accountId string) *Huobi {
	hb := new(Huobi)
	hb.baseUrl = "https://api.huobi.br.com"
	hb.httpClient = client
	hb.apikey = apikey
	hb.secretkey = secretkey
	hb.accountId = accountId
	hb.wsTickerHandleMap = make(map[string]func(*Ticker))
	return hb
}

// Close close fcoin websocket
func (hb *Huobi) Close() {
	if hb.ws != nil {
		hb.ws.Close()
	}
}

func (hb *Huobi) createWsConn() {
	if hb.ws != nil {
		return
	}

	hb.ws = NewWsConn("wss://api.huobi.br.com/ws")
	hb.ws.Heartbeat(func() interface{} {
		return map[string]interface{}{
			"ping": time.Now().Unix()}
	}, 5*time.Second)

	hb.ws.ReConnect()
	hb.ws.ReceiveMessage(func(msg []byte) {
		gzipreader, _ := gzip.NewReader(bytes.NewReader(msg))
		data, _ := ioutil.ReadAll(gzipreader)
		datamap := make(map[string]interface{})
		err := json.Unmarshal(data, &datamap)
		if err != nil {
			log.Println("json unmarshal error for ", string(data))
			return
		}

		if datamap["ping"] != nil {
			//log.Println(datamap)
			hb.ws.UpdateActivedTime()
			hb.ws.WriteJSON(map[string]interface{}{
				"pong": datamap["ping"]}) // pong
			return
		}

		if datamap["pong"] != nil { //
			hb.ws.UpdateActivedTime()
			return
		}

		if datamap["id"] != nil { //忽略订阅成功的回执消息
			log.Println(string(data))
			return
		}

		ch, isok := datamap["ch"].(string)
		if !isok {
			log.Println("error:", string(data))
			return
		}

		tick := datamap["tick"].(map[string]interface{})

		low, isok := tick["low"].(float64)
		if !isok {
			log.Println("error:", isok)
			return
		}
		ticker := &Ticker{Low: low}
		if hb.wsTickerHandleMap[ch] != nil {
			(hb.wsTickerHandleMap[ch])(ticker)
		}

	})
}

// GetTickerWithWs get ticker with websocket
func (hb *Huobi) GetTickerWithWs(pair CurrencyPair, handle func(ticker *Ticker)) error {
	hb.createWsConn()
	sub := fmt.Sprintf("market.%s.detail", strings.ToLower(pair.ToSymbol("")))
	hb.wsTickerHandleMap[sub] = handle
	return hb.ws.Subscribe(map[string]interface{}{
		"id":  1,
		"sub": sub})
}

func (hb *Huobi) GetExchangeName() string {
	return HUOBI
}
