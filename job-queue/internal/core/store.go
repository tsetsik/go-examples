package core

type (
	Store[A any] interface {
		Put(key string, item A) error
		Get(key string) (*A, error)
	}
)
