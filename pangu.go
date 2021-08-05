package schedule

import (
	`github.com/storezhang/pangu`
)

func init() {
	if err := pangu.New().Provides(newScheduler); nil != err {
		panic(err)
	}
}
