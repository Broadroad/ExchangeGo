package config

import (
	"fmt"
	"testing"
)

func TestNewConfig(t *testing.T) {
	c := NewConfig("./config.json")
	fmt.Println(c.Huobi.Frequence)
}
