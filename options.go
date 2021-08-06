package schedule

import (
	`time`
)

type options struct {
	scheduleType scheduleType

	cron string
	time time.Time

	id string
}

func defaultOptions() *options {
	return &options{}
}
