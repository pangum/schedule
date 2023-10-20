package plugin

import (
	"github.com/pangum/schedule/internal/config"
)

type Config struct {
	// 任务是否唯一
	Unique bool `json:"unique" yaml:"unique" xml:"unique" toml:"unique"`
	// 限制
	Limit *config.Limit `json:"limit" yaml:"limit" xml:"limit" toml:"limit"`
}
