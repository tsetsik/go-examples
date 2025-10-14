package store

import (
	"fmt"

	"github.com/tsetsik/go-examples/job-queue/internal/core"
)

type (
	cacheStore[A any] struct {
		data map[string]A
	}
)

func NewCacheStore[A any]() core.Store[A] {
	return &cacheStore[A]{
		data: make(map[string]A),
	}
}

func (s *cacheStore[A]) Put(key string, item A) error {
	s.data[key] = item
	return nil
}

func (s *cacheStore[A]) Get(key string) (*A, error) {
	item, exists := s.data[key]
	if !exists {
		return nil, fmt.Errorf("item with key %s not found", key)
	}

	return &item, nil
}
