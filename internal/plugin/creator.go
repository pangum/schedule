package plugin

import (
	"github.com/goexl/schedule"
	"github.com/goexl/simaqian"
	"github.com/pangum/pangu"
)

type Creator struct {
	// 解决命名空间问题
}

func (c *Creator) New(config *pangu.Config, logger simaqian.Logger) (scheduler *schedule.Scheduler, err error) {
	wrapper := new(Wrapper)
	if ge := config.Build().Get(wrapper); nil != ge {
		err = ge
	} else {
		scheduler, err = c.new(wrapper.Schedule, logger)
	}

	return
}

func (c *Creator) new(config *Config, logger simaqian.Logger) (scheduler *schedule.Scheduler, err error) {
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
