package plugin

import (
	"github.com/goexl/log"
	"github.com/goexl/schedule"
	"github.com/pangum/pangu"
)

type Constructor struct {
	// 构造方法
}

func (c *Constructor) New(config *pangu.Config, logger log.Logger) (scheduler *schedule.Scheduler, err error) {
	wrapper := new(Wrapper)
	if ge := config.Build().Get(wrapper); nil != ge {
		err = ge
	} else {
		scheduler, err = c.new(&wrapper.Schedule, logger)
	}

	return
}

func (c *Constructor) new(config *Config, logger log.Logger) (scheduler *schedule.Scheduler, err error) {
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
