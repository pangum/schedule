package schedule

type config struct {
	// 任务是否唯一
	Unique bool `json:"unique" yaml:"unique" xml:"unique" toml:"unique"`
}
