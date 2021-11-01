package schedule

import (
	`github.com/pangum/pangu`
)

func init() {
	if err := pangu.New().Provides(newScheduler); nil != err {
		panic(err)
	}
}
