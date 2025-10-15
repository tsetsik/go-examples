package core

type (
	ProcessedJobFunc func(job *Job)

	JobWorker interface {
		Start(fn ProcessedJobFunc)
		Enqueue(job *Job) error
		Stop()
	}
)
