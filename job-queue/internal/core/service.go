package core

import "fmt"

type (
	JobQueueService interface {
		Enqueue(job *Job) error
		Status(jobID string) (Status, error)
	}

	jobQueueService struct {
		jobWorker JobWorker
	}
)

func NewJobQueueService(jobWorker JobWorker) JobQueueService {
	return &jobQueueService{
		jobWorker: jobWorker,
	}
}

func (s *jobQueueService) Enqueue(job *Job) error {
	if err := job.Validate(); err != nil {
		return fmt.Errorf("invalid job: %w", err)
	}

	err := s.jobWorker.Enqueue(job)
	if err != nil {
		return fmt.Errorf("failed to enqueue job: %w", err)
	}

	return nil
}

func (s *jobQueueService) Status(jobID string) (Status, error) {
	// Implementation for checking job status
	return StatusQueued, nil
}
