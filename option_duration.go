package schedule

import (
	`time`
)

var _ option = (*optionDuration)(nil)

type optionDuration struct {
	duration time.Duration
}

// Duration 配置固定频率任务
func Duration(duration time.Duration) *optionDuration {
	return &optionDuration{duration: duration}
}

func (c *optionDuration) apply(options *options) {
	options.duration = c.duration
	options.scheduleType = scheduleTypeDuration
}
