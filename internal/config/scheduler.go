package config

import (
	"github.com/harluo/config"
)

type Scheduler struct {
	// 任务是否唯一
	Unique bool `json:"unique,omitempty"`
	// 限制
	Limit *Limit `json:"limit,omitempty"`
}

func newScheduler(config *config.Getter) (scheduler *Scheduler, err error) {
	scheduler = new(Scheduler)
	err = config.Get(&struct {
		Scheduler *Scheduler `json:"scheduler,omitempty" valschedulerate:"required"`
	}{
		Scheduler: scheduler,
	})

	return
}
