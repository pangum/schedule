package schedule

import (
	`time`
)

var _ option = (*optionDelay)(nil)

type optionDelay struct {
	delay time.Duration
}

// Delay 配置任务延迟执行
func Delay(delay time.Duration) *optionDelay {
	return &optionDelay{delay: delay}
}

func (c *optionDelay) apply(options *options) {
	options.delay = c.delay
}
