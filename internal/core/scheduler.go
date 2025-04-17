package core

import (
	"github.com/goexl/log"
	"github.com/goexl/schedule"
	"github.com/harluo/schedule/internal/config"
)

type Scheduler = schedule.Scheduler

func newScheduler(config *config.Scheduler, logger log.Logger) (scheduler *schedule.Scheduler, err error) {
	builder := schedule.New().Logger(logger)
	if config.Unique {
		builder.Unique()
	}

	// 全局限制条件
	limit := config.Limit
	if nil != limit {
		builder.Limit().Cpu(limit.Cpu).Memory(limit.Memory).Process(limit.Process).Max(limit.Max).Build()
	}
	scheduler = builder.Build()

	return
}
