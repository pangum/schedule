package schedule

import (
	`fmt`
	`math/rand`
	`time`
)

func fixTimeSpec(runTime time.Time, delayMaxRand int64, delay time.Duration) string {
	now := time.Now()
	if runTime.Before(now) {
		rand.Seed(time.Now().Unix())
		runTime = now.Add(time.Duration(rand.Int63n(delayMaxRand)) * time.Second)
	}
	if 0 != delay {
		runTime.Add(delay)
	}

	return fmt.Sprintf(
		"%d %d %d %d %d %d",
		runTime.Second(), runTime.Minute(), runTime.Hour(),
		runTime.Day(), runTime.Month(), runTime.Weekday(),
	)
}
