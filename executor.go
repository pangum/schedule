package schedule

type executor interface {
	// Run 执行
	Run() (err error)
}
