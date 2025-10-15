package core_test

import (
	"fmt"
	"log/slog"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/tsetsik/go-examples/job-queue/internal/core"
	"github.com/tsetsik/go-examples/job-queue/internal/core/mocks"
	"github.com/tsetsik/go-examples/job-queue/internal/store"
)

func TestService_Enqueue(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		workerMock := mocks.NewMockJobWorker(ctrl)

		job := &core.Job{
			ID:      "job1",
			Payload: "payload1",
			Status:  core.StatusQueued,
		}

		workerMock.EXPECT().Enqueue(job).Return(nil).Times(1)

		store := store.NewCacheStore[core.Job]()
		logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

		svc := core.NewJobQueueService(workerMock, store, logger)

		err := svc.Enqueue(job)

		require.NoError(t, err)

		storedJob, err := store.Get(job.ID)
		require.NoError(t, err)
		require.Equal(t, job, storedJob)
	})

	t.Run("invalid job", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		workerMock := mocks.NewMockJobWorker(ctrl)

		job := &core.Job{}

		workerMock.EXPECT().Enqueue(job).Return(nil).Times(0)

		store := store.NewCacheStore[core.Job]()
		logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

		svc := core.NewJobQueueService(workerMock, store, logger)

		err := svc.Enqueue(job)
		require.ErrorContains(t, err, "invalid job")
	})

	t.Run("enqueue error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		workerMock := mocks.NewMockJobWorker(ctrl)

		job := &core.Job{ID: "job1", Payload: "payload1"}

		jobEnqueueErr := fmt.Errorf("enqueue error")
		workerMock.EXPECT().Enqueue(job).Return(jobEnqueueErr).Times(1)

		store := store.NewCacheStore[core.Job]()
		logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

		svc := core.NewJobQueueService(workerMock, store, logger)

		err := svc.Enqueue(job)
		require.ErrorContains(t, err, "enqueue error")
	})
}

func TestService_Status(t *testing.T) {
	t.Parallel()

	t.Run("with-existing-job", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		workerMock := mocks.NewMockJobWorker(ctrl)
		store := store.NewCacheStore[core.Job]()
		logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

		svc := core.NewJobQueueService(workerMock, store, logger)

		job := &core.Job{ID: "job1", Payload: "payload1", Status: core.StatusQueued}
		err := store.Put(job.ID, *job)
		require.NoError(t, err)

		status, err := svc.Status(job.ID)
		require.NoError(t, err)
		require.Equal(t, core.StatusQueued, status)
	})
	t.Run("with-non-existing-job", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		workerMock := mocks.NewMockJobWorker(ctrl)
		store := store.NewCacheStore[core.Job]()
		logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

		svc := core.NewJobQueueService(workerMock, store, logger)

		status, err := svc.Status("non-existing-job")
		require.ErrorContains(t, err, "job not found")
		require.Equal(t, core.Status(""), status)
	})
}
