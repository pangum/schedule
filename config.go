package schedule

import (
	"time"
)

type config struct {
	// 任务是否唯一
	Unique bool `json:"unique" yaml:"unique" xml:"unique" toml:"unique"`
	// 延迟
	Delay time.Duration `json:"delay" yaml:"delay" xml:"delay" toml:"delay"`
}
