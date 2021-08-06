package schedule

type executor interface {
	Run() (err error)
}
