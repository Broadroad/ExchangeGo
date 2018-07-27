package main

import (
	"flag"
	"fmt"

	"github.com/ExchangeGo/scheduler"
)

const (
	configUsage = "use --config to define config path"
)

var (
	configPath = flag.String("config", "", configUsage)
)

func init() {
	flag.Parse()
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err is: ", err)
		}
	}()

	if len(*configPath) == 0 {
		panic(configUsage)
	}

	sc := scheduler.SchedulerConfig{Enablefc: true, Enablehuobi: false}
	scheduler := scheduler.NewScheduler(sc)
	scheduler.Schedule()
}
