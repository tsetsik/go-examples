package core

type (
	JobWorker interface {
		Start()
		Enqueue(job *Job) error
		Stop()
	}
)
