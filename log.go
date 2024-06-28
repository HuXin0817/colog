// Package colog provides a simple logging library with support for
// multiple log levels, colored console output, and concurrent-safe file writing.
package colog

import (
	"fmt"
	"os"
	"path"
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
	logFile *os.File
	mu      sync.Mutex
)

// OpenLog initializes the log file.
func OpenLog(filePath string) (err error) {
	if err = os.MkdirAll(path.Dir(filePath), os.ModePerm); err != nil {
		return err
	}
	if logFile, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		return err
	}
	return nil
}

// put writes a log message to the console and to the appropriate log file.
func put(msgType int, msg string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("\033[31m%v\033[0m\n", r) // Print panic message in red
		}
	}()

	// Get the caller's file and line number
	_, file, line, _ := runtime.Caller(3)
	t := time.Now()
	logMessage := fmt.Sprintf("%s %v\n", t.Format(timeFormat), msg)

	// Print the log message to the console with color formatting
	fmt.Printf("%s \033[35m%s\033[0m \033[34m\033[4m%s\033[0m line %d\033[0m: \"%v\"\n",
		head[msgType], t.Format(timeFormat), file, line, msg)

	// Write the log message to the log file
	if logFile != nil {
		mu.Lock()
		defer mu.Unlock()
		_, err := logFile.WriteString(logMessage)
		if err != nil {
			fmt.Printf("\033[31mFailed to write log to file: %v\033[0m\n", err)
		}
	}
}

// Error logs an error message.
func Error(msg ...any) {
	msgstr := ""
	for i, m := range msg {
		msgstr += fmt.Sprint(m)
		if i != len(msg)-1 {
			msgstr += " "
		}
	}
	put(errMsg, msgstr)
}

// Errorf logs a formatted error message.
func Errorf(format string, args ...any) {
	Error(fmt.Sprintf(format, args...))
}

// Info logs an info message.
func Info(msg ...any) {
	msgstr := ""
	for i, m := range msg {
		msgstr += fmt.Sprint(m)
		if i != len(msg)-1 {
			msgstr += " "
		}
	}
	put(infoMsg, msgstr)
}

// Infof logs a formatted info message.
func Infof(format string, args ...any) {
	Info(fmt.Sprintf(format, args...))
}

// Warn logs a warning message.
func Warn(msg ...any) {
	msgstr := ""
	for i, m := range msg {
		msgstr += fmt.Sprint(m)
		if i != len(msg)-1 {
			msgstr += " "
		}
	}
	put(warnMsg, msgstr)
}

// Warnf logs a formatted warning message.
func Warnf(format string, args ...any) {
	Warn(fmt.Sprintf(format, args...))
}
