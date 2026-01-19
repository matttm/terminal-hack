// Package logger provides centralized logging functionality for the terminal hacking game.
// All logs are written to a file for debugging and traceability.
package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	instance *Logger
	once     sync.Once
)

// Logger provides thread-safe logging to a file.
type Logger struct {
	file   *os.File
	logger *log.Logger
	mu     sync.Mutex
}

// Init initializes the global logger instance.
// Logs are written to terminal_hack.log in the current directory.
// Call this once at application startup.
func Init() error {
	var err error
	once.Do(func() {
		instance, err = newLogger()
	})
	return err
}

// newLogger creates a new logger instance with a timestamped log file.
func newLogger() (*Logger, error) {
	// Create logs directory if it doesn't exist
	logsDir := "./logs"
	if err := os.MkdirAll(logsDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create logs directory: %w", err)
	}

	// Create log file with timestamp
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	logFile := filepath.Join(logsDir, fmt.Sprintf("terminal_hack_%s.log", timestamp))

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	l := &Logger{
		file:   file,
		logger: log.New(file, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile),
	}

	l.log("INFO", "Logger initialized", "logFile", logFile)
	return l, nil
}

// Close closes the log file. Call this before application exit.
func Close() error {
	if instance != nil && instance.file != nil {
		return instance.file.Close()
	}
	return nil
}

// Info logs an informational message with optional key-value pairs.
func Info(msg string, keysAndValues ...interface{}) {
	if instance != nil {
		instance.log("INFO", msg, keysAndValues...)
	}
}

// Debug logs a debug message with optional key-value pairs.
func Debug(msg string, keysAndValues ...interface{}) {
	if instance != nil {
		instance.log("DEBUG", msg, keysAndValues...)
	}
}

// Error logs an error message with optional key-value pairs.
func Error(msg string, keysAndValues ...interface{}) {
	if instance != nil {
		instance.log("ERROR", msg, keysAndValues...)
	}
}

// Warn logs a warning message with optional key-value pairs.
func Warn(msg string, keysAndValues ...interface{}) {
	if instance != nil {
		instance.log("WARN", msg, keysAndValues...)
	}
}

// log is the internal logging method that formats and writes log entries.
func (l *Logger) log(level, msg string, keysAndValues ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Format key-value pairs
	kvStr := ""
	for i := 0; i < len(keysAndValues); i += 2 {
		if i+1 < len(keysAndValues) {
			kvStr += fmt.Sprintf(" %v=%v", keysAndValues[i], keysAndValues[i+1])
		}
	}

	// Write log entry
	l.logger.Output(3, fmt.Sprintf("[%s] %s%s", level, msg, kvStr))
}
