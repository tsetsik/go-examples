package processor

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	existingFiles = []string{
		"../assets/example1.json",
	}
)

func Test_Process(t *testing.T) {
	t.Parallel()

	t.Run("process-existing-files", func(t *testing.T) {
		t.Parallel()

		logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		processor := NewFilesProcessor(logger, 3)

		filesPath := getTestFilesPath(t)

		got := processor.Process(filesPath)

		require.NotNil(t, logger)
		require.NotNil(t, got)
	})
}

func getTestFilesPath(t *testing.T) []string {
	t.Helper()

	var filesPath []string
	for _, file := range existingFiles {
		ex, err := os.Executable()
		require.NoError(t, err)

		exPath := fmt.Sprintf("%s/../%s", filepath.Dir(ex), file)
		filesPath = append(filesPath, exPath)
	}

	return filesPath
}
