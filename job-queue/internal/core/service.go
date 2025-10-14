package core

import (
	"fmt"
	"sync"
)

type (
	JobQueueService interface {
		Enqueue(job *Job) error
		Status(jobID string) (Status, error)
	}

	jobQueueService struct {
		jobWorker JobWorker
		m         sync.RWMutex
		store     Store[Job]
	}
)

func NewJobQueueService(jobWorker JobWorker, store Store[Job]) JobQueueService {
	return &jobQueueService{
		jobWorker: jobWorker,
		store:     store,
	}
}

func (s *jobQueueService) Enqueue(job *Job) error {
	s.m.Lock()
	defer s.m.Unlock()

	if err := job.Validate(); err != nil {
		return fmt.Errorf("invalid job: %w", err)
	}

	job.Status = StatusQueued

	err := s.jobWorker.Enqueue(job)
	if err != nil {
		return fmt.Errorf("failed to enqueue job: %w", err)
	}

	err = s.store.Put(job.ID, *job)
	if err != nil {
		return fmt.Errorf("failed to store job: %w", err)
	}

	return nil
}

func (s *jobQueueService) Status(jobID string) (Status, error) {
	s.m.RLock()
	defer s.m.RUnlock()

	job, err := s.store.Get(jobID)
	if err != nil {
		return "", fmt.Errorf("failed to get job: %w", err)
	}

	if job == nil {
		return "", fmt.Errorf("job not found")
	}

	return job.Status, nil
}
