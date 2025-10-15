package core

import (
	"fmt"

	validator "github.com/go-playground/validator/v10"
)

type (
	Status string

	Job struct {
		ID      string `json:"job_id" validate:"required"`
		Payload any    `json:"payload" validate:"required"`
		Status  Status `json:"status"`
	}

	JobProcessed struct {
		Job  Job
		Code int
		Err  error
	}
)

var (
	StatusQueued Status = "QUEUED"
	StatusDone   Status = "DONE"
)

func (j *Job) Validate() error {
	if j == nil {
		return fmt.Errorf("job is nil")
	}

	validator := validator.New()
	return validator.Struct(j)
}
