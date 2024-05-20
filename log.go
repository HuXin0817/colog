// Package colog provides a simple logging library with support for
// multiple log levels, colored console output, and concurrent-safe file writing.
package colog

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

// Constants for log message types
const (
	errMsg  = iota // Error message
	infoMsg        // Info message
	warnMsg        // Warning message
)

// Time format used in log entries
const timeFormat = "2006-01-02 15:04:05"

// Global variables
var (
	head = []string{
		errMsg:  "\033[33m[ERROR]\033[0m", // Yellow for error messages
		infoMsg: "\033[32m[INFO]\033[0m",  // Green for info messages
		warnMsg: "\033[31m[WARN]\033[0m",  // Red for warning messages
	}
	record bool             // Indicates if logging to files is enabled
	files  map[int]*os.File // Map of file handles for different log levels
	locks  []*sync.Mutex    // Mutex locks for concurrent-safe file writing
)

// Open initializes log files in the specified directory.
// It creates separate files for error, info, and warning messages.
func Open(dirPath string) (err error) {
	if record {
		return nil // Already initialized
	}

	files = make(map[int]*os.File)
	locks = []*sync.Mutex{
		errMsg:  {},
		infoMsg: {},
		warnMsg: {},
	}

	// Create the directory if it doesn't exist
	if err = os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return err
	}

	// Log file names for different log levels
	logFileName := []string{
		errMsg:  "error.log",
		infoMsg: "info.log",
		warnMsg: "warn.log",
	}

	// Open log files for appending
	for i, file := range logFileName {
		if files[i], err = os.OpenFile(filepath.Join(dirPath, file), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
			return err
		}
	}

	record = true // Mark logging to files as enabled
	return nil
}

// put writes a log message to the console and to the appropriate log file.
func put(msgType int, msg any) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("\033[31m%v\033[0m\n", r) // Print panic message in red
		}
	}()

	// Get the caller's file and line number
	_, file, line, _ := runtime.Caller(3)
	t := time.Now()

	// Print log message to the console with color formatting
	fmt.Printf("%s \033[35m%s\033[0m \033[34m\033[4m%s\033[0m line %d\033[0m: \"%v\"\n",
		head[msgType], t.Format(timeFormat), file, line, msg)

	// Write log message to the file if recording is enabled
	if record {
		locks[msgType].Lock() // Acquire lock for the log level
		mess := fmt.Sprintf("%s %s line %d: %v\n", t.Format(timeFormat), file, line, msg)
		if _, err := files[msgType].WriteString(mess); err != nil {
			fmt.Printf("\033[31m%v\033[0m\n", err) // Print error message in red
		}
		locks[msgType].Unlock() // Release lock for the log level
	}
}

// Error logs an error message.
func Error(msg any) { put(errMsg, msg) }

// Errorf logs a formatted error message.
func Errorf(format string, args ...any) { Error(fmt.Sprintf(format, args...)) }

// Info logs an info message.
func Info(msg any) { put(infoMsg, msg) }

// Infof logs a formatted info message.
func Infof(format string, args ...any) { Info(fmt.Sprintf(format, args...)) }

// Warn logs a warning message.
func Warn(msg any) { put(warnMsg, msg) }

// Warnf logs a formatted warning message.
func Warnf(format string, args ...any) { Warn(fmt.Sprintf(format, args...)) }
