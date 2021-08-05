package schedule

import (
	`github.com/robfig/cron/v3`
)

// Scheduler 任务计划组织程序
type Scheduler struct {
	cron    *cron.Cron
	started bool
	stopped bool
}

func newScheduler() *Scheduler {
	return &Scheduler{
		cron:    cron.New(),
		started: false,
		stopped: false,
	}
}

func (s *Scheduler) Add(executor executor, opts ...option) {

}

func (s *Scheduler) Start(opts ...option) {
	if !s.started {
		s.cron.Start()
		s.started = true
	}
}

func (s *Scheduler) Stop(opts ...option) {
	if !s.stopped {
		s.cron.Stop()
		s.stopped = true
	}
}
