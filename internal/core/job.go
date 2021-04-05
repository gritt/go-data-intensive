package core

const (
	JOB_PENDING = iota
	JOB_DONE
	JOB_ERROR
)

type Job struct {
	Data   string
	UUID   string
	Status int
}
