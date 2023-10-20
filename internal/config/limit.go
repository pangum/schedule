package config

type Limit struct {
	// 最大使用百分比
	Cpu float64 `json:"cpu" yaml:"cpu" xml:"cpu" toml:"cpu"`
	// 最大内存占用百分比
	Memory float64 `json:"memory" yaml:"memory" xml:"memory" toml:"memory"`
	// 最大进程数量
	Process int `json:"process" yaml:"process" xml:"process" toml:"process"`
	// 最大任务数量
	Max int `json:"max" yaml:"max" xml:"max" toml:"max"`
}
