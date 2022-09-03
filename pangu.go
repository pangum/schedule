package schedule

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Dependencies(
		newScheduler,
	)
}
