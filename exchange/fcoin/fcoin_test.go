package fcoin

import (
	"log"
	"net/http"
	"testing"
)

func TestNewFCoin(t *testing.T) {
	fc := NewFCoin(&http.Client{}, "", "")
	log.Println(fc.GetExchangeName(), fc.GetTimeOffset())
}
