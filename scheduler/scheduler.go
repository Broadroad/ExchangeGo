package scheduler

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ExchangeGo/exchange/fcoin"
	"github.com/ExchangeGo/exchange/huobi"
)

type SchedulerConfig struct {
	enablefc    bool
	enablehuobi bool
}

type scheduler struct {
	sc    SchedulerConfig
	fc    *fcoin.FCoin
	huobi *huobi.Huobi
}

func NewScheduler(sc SchedulerConfig) *scheduler {
	s := &scheduler{sc: sc}
	if s.enablefc {
		s.fc = fcoin.NewFCoin(http.DefaultClient, "", "")
	}
	if s.enablehuobi {
		s.huobi = huobi.NewHuobi(http.DefaultClient, "", "", "")
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
		if s.fc {
			s.fc.Close()
		}
		if s.huobi {
			s.huobi.Close()
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
	<-done
	fmt.Println("exiting")

}
