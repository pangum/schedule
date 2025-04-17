package config

type Limit struct {
	// 最大使用百分比
	Cpu float64 `json:"cpu,omitempty"`
	// 最大内存占用百分比
	Memory float64 `json:"memory,omitempty"`
	// 最大进程数量
	Process int `json:"process,omitempty"`
	// 最大任务数量
	Max int `json:"max,omitempty"`
}
