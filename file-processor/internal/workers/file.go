package workers

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"

	"github.com/tsetsik/go-examples/file-processor/internal/core"
)

func (wp *workerPool) ProcessFileWorker(files <-chan string, results chan<- core.FileProcessed) {
	for file := range files {
		wp.logger.Info("Worker started processing file", slog.String("file", file))

		f, err := os.Open(file)
		if err != nil {
			results <- core.FileProcessed{Path: file, Err: fmt.Errorf("failed to open file: %w", err)}
			continue
		}

		scanner := bufio.NewScanner(f)
		lines := 0
		for scanner.Scan() {
			lines++
		}

		if err := f.Close(); err != nil {
			results <- core.FileProcessed{Path: file, Err: fmt.Errorf("failed to close file: %w", err)}
			continue
		}

		if err := scanner.Err(); err != nil {
			results <- core.FileProcessed{Path: file, Err: fmt.Errorf("failed to read file: %w", err)}
			continue
		}

		results <- core.FileProcessed{Path: file, Lines: lines}
	}
}
