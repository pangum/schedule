package schedule

type executor interface {
	run() (err error)
}
