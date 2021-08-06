package schedule

const (
	scheduleTypeCron     scheduleType = "cron"
	scheduleTypeDuration scheduleType = "duration"
	scheduleTypeTime     scheduleType = "time"
)

type scheduleType string
