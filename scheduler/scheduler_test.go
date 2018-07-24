package scheduler

import (
	"net/http"
	"testing"

	"github.com/ExchangeGo/exchange/fcoin"
)

func TestNewFCoin(t *testing.T) {
	fc := fcoin.NewFCoin(&http.Client{}, "", "")
	sconfig := SchedulerConfig{fc: fc}
	NewScheduler(sconfig)
}
