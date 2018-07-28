package scheduler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ExchangeGo/common"
	"github.com/ExchangeGo/exchange/fcoin"
	"github.com/ExchangeGo/exchange/huobi"
)

type SchedulerConfig struct {
	Enablefc    bool
	Enablehuobi bool
}

type scheduler struct {
	sc SchedulerConfig
	fc *fcoin.FCoin
	hb *huobi.Huobi
}

func NewScheduler(sc SchedulerConfig) *scheduler {
	s := &scheduler{sc: sc}
	if sc.Enablefc {
		s.fc = fcoin.NewFCoin(http.DefaultClient, "", "")
	}
	if sc.Enablehuobi {
		s.hb = huobi.NewHuobi(http.DefaultClient, "", "", "")
	}
	return s
}

// Schedule schedule the cralwe tasks
func (s *scheduler) Schedule() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		// close all the websocket
		if s.fc != nil {
			s.fc.Close()
		}
		if s.hb != nil {
			s.hb.Close()
		}
	}()

	go func() {
		sig := <-sigs
		fmt.Println("recieve signal: ", sig)

		done <- true
	}()

	// The program will wait here until it gets the
	// expected signal (as indicated by the goroutine
	// above sending a value on `done`) and then exit.
	fmt.Println("awaiting signal")

	go s.schedule()

	<-done
	fmt.Println("exiting")

}

func (s *scheduler) schedule() {
	if s.sc.Enablehuobi {
		s.hb.GetTickerWithWs(common.BTC_USDT, func(ticker *common.Ticker) {
			log.Println(ticker)
		})
	}
}
