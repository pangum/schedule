package schedule

import (
	`sync`

	`github.com/robfig/cron/v3`
	`github.com/rs/xid`
)

// Scheduler 任务计划组织程序
type Scheduler struct {
	cron    *cron.Cron
	idCache sync.Map

	started bool
	stopped bool
}

func newScheduler() *Scheduler {
	return &Scheduler{
		cron:    cron.New(),
		idCache: sync.Map{},

		started: false,
		stopped: false,
	}
}

func (s *Scheduler) Add(executor executor, opts ...option) (id string, err error) {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}

	if "" == options.id {
		id = xid.New().String()
	}

	var entryId cron.EntryID
	switch options.scheduleType {
	case scheduleTypeCron:
		entryId, err = s.cron.AddFunc(options.cron, func() {
			executor.run()
		})
	}

	return
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
