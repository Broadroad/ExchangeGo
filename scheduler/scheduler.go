package scheduler

type struct schedulerConfig {
	fc	*fcoin.FCoin
	huobi *huobi.Huobi
}

type struct scheduler {
	sc schedulerConfig
}


func NewScheduler(sc schedulerConfig) *scheduler {
	s := &scheduler{sc: sc}
	return s
}

// Schedule schedule the cralwe tasks
func (s *scheduler)Schedule() {
	
}
