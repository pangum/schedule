package schedule

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/robfig/cron/v3"
	"github.com/rs/xid"
)

// Scheduler 任务计划组织程序
type Scheduler struct {
	cron *cron.Cron
	ids  *sync.Map

	started bool
	stopped bool
}

func newScheduler() (scheduler *Scheduler) {
	scheduler = &Scheduler{
		cron: cron.New(cron.WithSeconds()),
		ids:  new(sync.Map),

		started: false,
		stopped: false,
	}
	scheduler.Start()

	return
}

func (s *Scheduler) Add(executor executor, opts ...option) (id string, err error) {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	if `` == _options.id {
		id = xid.New().String()
	}

	var entryId cron.EntryID
	switch _options.scheduleType {
	case scheduleTypeCron:
		entryId, err = s.cron.AddFunc(_options.cron, func() {
			_ = executor.Run()
		})
	case scheduleTypeDuration:
		entryId, err = s.cron.AddFunc(fmt.Sprintf("@every %s", _options.duration.String()), func() {
			_ = executor.Run()
		})
	case scheduleTypeTime:
		entryId, err = s.cron.AddFunc(fixTimeSpec(_options.time, _options.delayMaxRand, _options.delay), func() {
			_ = executor.Run()
		})
	}
	if nil != err {
		return
	}
	id = strconv.Itoa(int(entryId))
	s.ids.Store(id, entryId)

	return
}

func (s *Scheduler) Start(_ ...option) {
	if !s.started {
		s.cron.Start()
		s.started = true
	}
}

func (s *Scheduler) Stop(_ ...option) {
	if !s.stopped {
		s.cron.Stop()
		s.stopped = true
	}
}

func (s *Scheduler) Remove(id *optionId) {
	if _id, ok := s.ids.Load(id.id); ok {
		s.cron.Remove(_id.(cron.EntryID))
		s.ids.Delete(id.id)
	}
}

func (s *Scheduler) Clear() {
	for _, entry := range s.cron.Entries() {
		s.cron.Remove(entry.ID)
	}
	s.ids = new(sync.Map)
}

func (s *Scheduler) Contains(id *optionId) (contains bool) {
	_, contains = s.ids.Load(id.id)

	return
}
