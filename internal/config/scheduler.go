package config

import (
	"github.com/harluo/boot"
)

type Scheduler struct {
	// 任务是否唯一
	Unique bool `json:"unique,omitempty"`
	// 限制
	Limit *Limit `json:"limit,omitempty"`
}

func newScheduler(config *boot.Config) (scheduler *Scheduler, err error) {
	scheduler = new(Scheduler)
	err = config.Build().Get(&struct {
		Scheduler *Scheduler `json:"scheduler,omitempty" valschedulerate:"required"`
	}{
		Scheduler: scheduler,
	})

	return
}
