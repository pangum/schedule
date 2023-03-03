package schedule

import (
	"github.com/goexl/schedule"
	"github.com/pangum/logging"
	"github.com/pangum/pangu"
)

func newScheduler(config *pangu.Config, logger *logging.Logger) (scheduler *Scheduler, err error) {
	conf := new(panguConfig)
	if err = config.Load(conf); nil != err {
		return
	}

	_schedule := conf.Schedule
	builder := schedule.New().Logger(logger.Logger)
	if _schedule.Unique {
		builder.Unique()
	}
	scheduler = builder.Build()

	return
}
