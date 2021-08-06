package schedule

var _ option = (*optionCron)(nil)

type optionCron struct {
	cron string
}

// Cron 配置Cron表达式任务
func Cron(cron string) *optionCron {
	return &optionCron{cron: cron}
}

func (c *optionCron) apply(options *options) {
	options.cron = c.cron
	options.scheduleType = scheduleTypeCron
}
