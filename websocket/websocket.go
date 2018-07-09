package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type WsConn struct {
	*websocket.Conn
	url                      string
	heartbeatIntervalTime    time.Duration
	checkConnectIntervalTime time.Duration
	actived                  time.Time
	close                    chan int
	isClose                  bool
	subs                     []interface{}
}

const (
	SUB_TICKER      = 1 + iota
	SUB_ORDERBOOK
	SUB_KLINE_1M
	SUB_KLINE_15M
	SUB_KLINE_30M
	SUB_KLINE_1D
	UNSUB_TICKER
	UNSUB_ORDERBOOK
)

func NewWsConn(wsurl string) *WsConn {
	wsConn, _, err := websocket.DefaultDialer.Dial(wsurl, nil)
	if err != nil {
		panic(err)
	}
	return &WsConn{Conn: wsConn, url: wsurl, actived: time.Now(), checkConnectIntervalTime: 30 * time.Second, close: make(chan int, 1)}
}

// ReConnect check the connect every checkConnectIntervalTime
func (ws *WsConn) ReConnect() {

	timer := time.NewTimer(ws.checkConnectIntervalTime)
	go func() {
		for {
			select {
			case <-timer.C:
				if time.Now().Sub(ws.actived) >= 2*ws.checkConnectIntervalTime {
					ws.Close()
					log.Println("start reconnect websocket:", ws.url)
					wsConn, _, err := websocket.DefaultDialer.Dial(ws.url, nil)
					if err != nil {
						log.Println("reconnect fail ???")
					} else {
						ws.Conn = wsConn
						ws.actived = time.Now()
						//re subscribe
						for _, sub := range ws.subs {
							log.Println("subscribe:", sub)
							ws.WriteJSON(sub)
						}
					}
				}
				timer.Reset(ws.checkConnectIntervalTime)
			case <-ws.close:
				timer.Stop()
				log.Println("close websocket connect, exiting reconnect goroutine.")
				return
			}
		}
	}()
}