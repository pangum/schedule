package schedule

type option interface {
	apply(options *options)
}
