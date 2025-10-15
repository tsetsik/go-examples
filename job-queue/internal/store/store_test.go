package store

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Put(t *testing.T) {
	t.Parallel()

	t.Run("put-and-get-item", func(t *testing.T) {
		t.Parallel()

		store := NewCacheStore[int]()

		err := store.Put("key1", 42)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		item, err := store.Get("key1")
		require.NoError(t, err)
		require.Equal(t, 42, *item)
	})

	t.Run("get-non-existing-item", func(t *testing.T) {
		t.Parallel()

		store := NewCacheStore[int]()

		item, err := store.Get("non-existing-key")
		require.ErrorContains(t, err, "item with key non-existing-key not found")
		require.Nil(t, item)
	})
}
