package core

type (
	FileProcessed struct {
		Path  string `json:"path"`
		Lines int    `json:"lines"`
		Err   error  `json:"error,omitempty"`
	}
)
