package logger

import (
	"log"
	"log/slog"
	"os"

	slogmulti "github.com/samber/slog-multi"
)

func Start(filePath string) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	fileHandler := slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})
	consoleHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	multi := slog.New(slogmulti.Fanout(consoleHandler, fileHandler))
	slog.SetDefault(multi)
}
