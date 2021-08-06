package schedule

import (
	`time`
)

var _ option = (*optionTime)(nil)

type optionTime struct {
	time time.Time
}

// Time 配置指定时间任务
func Time(time time.Time) *optionTime {
	return &optionTime{time: time}
}

// DurationTime 配置指定时间任务
func DurationTime(duration time.Duration) *optionTime {
	return &optionTime{time: time.Now().Add(duration)}
}

func (c *optionTime) apply(options *options) {
	options.time = c.time
	options.scheduleType = scheduleTypeTime
}
