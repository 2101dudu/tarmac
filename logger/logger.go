package logger

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

// not in use
const (
	INFO = iota
	DEBUG
	TIMESTAMP
)

//

type Logger struct {
	file     *os.File
	writer   *bufio.Writer
	toStdout bool
}

var Log *Logger // global pointer

func Start(filePath string, toStdout bool) {
	file, err := os.OpenFile(filePath, os.O_APPEND, 0777) // or os.OpenFile() with O_APPEND for long-running logs
	if err != nil {
		log.Fatal(err)
	}

	Log = &Logger{
		file:     file,
		writer:   bufio.NewWriter(file),
		toStdout: toStdout,
	}
}

func (l *Logger) Log(args ...any) {
	timestamp := time.Now().Format(time.RFC3339)
	logLine := fmt.Sprintf("[%s] %s\n", timestamp, fmt.Sprint(args...))

	if l.toStdout {
		fmt.Print(logLine)
	}

	l.writer.WriteString(logLine) // buffer it
}

func (l *Logger) Close() error {
	err := l.writer.Flush()
	if err != nil {
		return err
	}
	return l.file.Close()
}

func (l *Logger) TrackTime() func() {
	start := time.Now()
	return func() {
		duration := time.Since(start)
		if duration.Milliseconds() < 5 {
			return
		}
		visual := strings.Repeat("=", int(duration.Milliseconds()/1)) // 1ms per bar
		msg := fmt.Sprintf("	[TIMER] %s took %s |%s>", callerName(), duration, visual)

		if l.toStdout {
			fmt.Println(msg)
		}
		l.writer.WriteString(msg + "\n")
	}
}

func callerName() string {
	pc, _, _, ok := runtime.Caller(2) // skip this function call and logger function call
	if !ok {
		return ""
	}
	f := runtime.FuncForPC(pc)
	if f == nil {
		return ""
	}
	return f.Name()
}
