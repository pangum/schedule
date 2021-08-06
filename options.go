package schedule

import (
	`time`
)

type options struct {
	scheduleType scheduleType

	cron     string
	time     time.Time
	delay    int64
	duration time.Duration

	id string
}

func defaultOptions() *options {
	return &options{
		delay: 10,
	}
}
