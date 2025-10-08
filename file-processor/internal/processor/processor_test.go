package processor

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Process(t *testing.T) {
	t.Parallel()

	t.Run("process-existing-files", func(t *testing.T) {
		t.Parallel()

		logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		processor := NewFilesProcessor(logger, 3)
		got := processor.Process()

		require.NotNil(t, logger)
	})
}
