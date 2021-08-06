package schedule

import (
	`fmt`
	`math/rand`
	`time`
)

func fixTimeSpec(runTime time.Time, delay int64) string {
	now := time.Now()
	if runTime.Before(now) {
		rand.Seed(time.Now().Unix())
		runTime = now.Add(time.Duration(rand.Int63n(delay)) * time.Second)
	}

	return fmt.Sprintf(
		"%d %d %d %d %d %d",
		runTime.Second(), runTime.Minute(), runTime.Hour(),
		runTime.Day(), runTime.Month(), runTime.Weekday(),
	)
}
