package schedule

type config struct {
	// 任务是否唯一
	Unique bool `json:"unique" yaml:"unique" xml:"unique" toml:"unique"`
	// 限制
	Limit *limit `json:"limit" yaml:"limit" xml:"limit" toml:"limit"`
}
