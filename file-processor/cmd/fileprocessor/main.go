package main

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"

	"github.com/tsetsik/go-examples/file-processor/internal/processor"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
	processor := processor.NewFilesProcessor(logger, 3)

	_, b, _, _ := runtime.Caller(0)
	path := filepath.Dir(fmt.Sprintf("%s/../../../assets", b))

	output := processor.Process([]string{
		fmt.Sprintf("%s/%s", path, "/assets/example1.json"),
		fmt.Sprintf("%s/%s", path, "/assets/example2.json"),
	})

	for _, fileOutput := range output {
		if fileOutput.Err == nil {
			logger.Info("File processed successfully", slog.String("file", fileOutput.Path), slog.Int("lines", fileOutput.Lines))
		} else {
			logger.Error("File processing failed", slog.String("file", fileOutput.Path), slog.String("error", fileOutput.Err.Error()))
		}
	}

	logger.Debug("Processing output", slog.Any("output", output))
}
