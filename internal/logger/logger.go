package logger

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

// Logger wraps the standard library logger with additional methods.
type Logger struct {
	logger *log.Logger
}

// NewLogger sets up a logger with timestamp and additional settings.
func NewLogger() *Logger {
	return &Logger{
		// logger: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile),
		logger: log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
	}
}

// Info logs a message at Info level.
func (l *Logger) Info(v ...interface{}) {
	l.log("INFO", v...)
}

// Infof logs a formatted message at Info level.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.logf("INFO", format, v...)
}

// Error logs a message at Error level.
func (l *Logger) Error(v ...interface{}) {
	l.log("ERROR", v...)
}

// Errorf logs a formatted message at Error level.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.logf("ERROR", format, v...)
}

// Warn logs a message at Warn level.
func (l *Logger) Warn(v ...interface{}) {
	l.log("WARN", v...)
}

// Warnf logs a formatted message at Warn level.
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.logf("WARN", format, v...)
}

// Fatal logs a message and exits the application.
func (l *Logger) Fatal(v ...interface{}) {
	l.log("FATAL", v...)
	os.Exit(1)
}

// Fatalf logs a formatted message and exits the application.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.logf("FATAL", format, v...)
	os.Exit(1)
}

// WithField adds a field to the logger.
func (l *Logger) WithField(key string, value interface{}) *Logger {
	return &Logger{
		logger: l.logger,
	}
}

// log logs a message with the specified level.
func (l *Logger) log(level string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	files := strings.Split(file, "/")
	if len(files) > 2 {
		file = "/" + strings.Join(files[len(files)-2:], "/")
	} else {
		file = "/" + strings.Join(files[len(files)-1:], "/")
	}

	prefix := fmt.Sprintf("%s %s:%d ", level, file, line)
	l.logger.SetPrefix(prefix)
	l.logger.Println(v...)
}

// logf logs a formatted message with the specified level.
func (l *Logger) logf(level, format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	files := strings.Split(file, "/")
	if len(files) > 2 {
		file = "/" + strings.Join(files[len(files)-2:], "/")
	} else {
		file = "/" + strings.Join(files[len(files)-1:], "/")
	}

	prefix := fmt.Sprintf("%s %s:%d ", level, file, line)
	l.logger.SetPrefix(prefix)
	l.logger.Printf(format, v...)
}

func (l *Logger) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Capture status code
		rr := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(rr, r)

		duration := time.Since(start)

		log.Printf("[%s] %s %d %s", r.Method, r.URL.Path, rr.statusCode, duration)
	})
}

// Helper to capture status code
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rr *responseRecorder) WriteHeader(code int) {
	rr.statusCode = code
	rr.ResponseWriter.WriteHeader(code)
}