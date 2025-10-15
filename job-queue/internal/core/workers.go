package core

//go:generate mockgen -source=./workers.go -destination=./mocks/mock_workers.go -package=mocks

type (
	ProcessedJobFunc func(job *Job)

	JobWorker interface {
		Start(fn ProcessedJobFunc)
		Enqueue(job *Job) error
		Stop()
	}
)
