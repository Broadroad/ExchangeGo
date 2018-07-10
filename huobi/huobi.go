package huobi

import (
	"github.com/ExchangeGo/websocket"
)

// Huobi is the struct to buy/sell/subscribe
type Huobi struct {
	ws *websocket.WsConn
}


