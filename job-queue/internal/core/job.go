package core

import (
	"fmt"

	"github.com/go-playground/validator"
)

type (
	Status string

	Job struct {
		ID      string `json:"job_id" required:"true"`
		Payload any    `json:"payload" required:"true"`
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
