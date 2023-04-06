package schedule

import (
	"github.com/goexl/schedule"
	"github.com/pangum/logging"
	"github.com/pangum/pangu"
)

func newScheduler(config *pangu.Config, logger logging.Logger) (scheduler *Scheduler, err error) {
	conf := new(wrapper)
	if err = config.Load(conf); nil != err {
		return
	}

	wrap := conf.Schedule
	builder := schedule.New().Logger(logger)
	if wrap.Unique {
		builder.Unique()
	}

	// 全局限制条件
	_limit := wrap.Limit
	if nil != _limit {
		builder.Limit().Cpu(_limit.Cpu).Memory(_limit.Memory).Process(_limit.Process).Max(_limit.Max).Build()
	}
	scheduler = builder.Build()

	return
}
