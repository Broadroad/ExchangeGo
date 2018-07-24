package scheduler

import (
	"github.com/ExchangeGo/exchange/fcoin"
	"github.com/ExchangeGo/exchange/huobi"
)

type SchedulerConfig struct {
	fc    *fcoin.FCoin
	huobi *huobi.Huobi
}

type scheduler struct {
	sc SchedulerConfig
}

func NewScheduler(sc SchedulerConfig) *scheduler {
	s := &scheduler{sc: sc}
	return s
}

// Schedule schedule the cralwe tasks
func (s *scheduler) Schedule() {
}
