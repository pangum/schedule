package schedule

import (
	`time`
)

type options struct {
	scheduleType scheduleType

	cron         string
	time         time.Time
	delayMaxRand int64
	delay        time.Duration
	duration     time.Duration

	id string
}

func defaultOptions() *options {
	return &options{
		delayMaxRand: 10,
	}
}
