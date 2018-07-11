package common

import (
	"strings"
	"testing"
)

func TestCurrency2String(t *testing.T) {
	btc := NewCurrency("btc", "bitcoin")
	t.Log(strings.ToUpper(btc.String()))
}
