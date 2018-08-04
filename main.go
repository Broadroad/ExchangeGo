package main

import (
	"flag"
	"fmt"

	"github.com/ExchangeGo/config"
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

	fmt.Println(*configPath)

	var sc scheduler.SchedulerConfig
	fmt.Println("hello ", sc)
	if len(*configPath) == 0 {
		panic(configUsage)
	} else {
		config := config.NewConfig(*configPath)
		// Parse config to Scheduler config
		if config.Huobi != nil {
			sc.Enablehuobi = true
			sc.HuobiConfig = config.Huobi
		}
		if config.FCoin != nil {
			sc.Enablefc = true
			sc.FCoinConfig = config.FCoin
		}
	}

	scheduler := scheduler.NewScheduler(sc)
	scheduler.Schedule()
}
