package main

/*
#cgo LDFLAGS: -llog

#include <android/log.h>
#include <stdlib.h>

void android_log(int prio, const char* tag, const char* msg) {
    __android_log_write(prio, tag, msg);
}
*/
import "C"
import (
	"bytes"
	"fmt"
	"unsafe"

	"golang.org/x/sys/unix"
)

type androidLogger struct {
	enabled  bool
	severity Severity
}

func (self androidLogger) androidLog(level C.int, tag string, message string) {
	C.android_log(level, cstring(tag), cstring(message))
}

func (self androidLogger) Verbose(tag string, message string) {
	if self.enabled && self.severity <= C.ANDROID_LOG_VERBOSE {
		self.androidLog(C.ANDROID_LOG_VERBOSE, tag, message)
	}
}

func (self androidLogger) Verbosef(tag string, format string, args ...interface{}) {
	self.Verbose(tag, fmt.Sprintf(format, args...))
}

func (self androidLogger) Debug(tag string, message string) {
	if self.enabled && self.severity <= C.ANDROID_LOG_DEBUG {
		self.androidLog(C.ANDROID_LOG_DEBUG, tag, message)
	}
}

func (self androidLogger) Debugf(tag string, format string, args ...interface{}) {
	self.Debug(tag, fmt.Sprintf(format, args...))
}

func (self androidLogger) Info(tag string, message string) {
	if self.enabled && self.severity <= C.ANDROID_LOG_INFO {
		self.androidLog(C.ANDROID_LOG_INFO, tag, message)
	}
}

func (self androidLogger) Infof(tag string, format string, args ...interface{}) {
	self.Info(tag, fmt.Sprintf(format, args...))
}

func (self androidLogger) Warn(tag string, message string) {
	if self.enabled && self.severity <= C.ANDROID_LOG_WARN {
		self.androidLog(C.ANDROID_LOG_WARN, tag, message)
	}
}

func (self androidLogger) Warnf(tag string, format string, args ...interface{}) {
	self.Warn(tag, fmt.Sprintf(format, args...))
}

func (self androidLogger) Error(tag string, message string) {
	if self.enabled && self.severity <= C.ANDROID_LOG_ERROR {
		self.androidLog(C.ANDROID_LOG_ERROR, tag, message)
	}
}

func (self androidLogger) Errorf(tag string, format string, args ...interface{}) {
	self.Error(tag, fmt.Sprintf(format, args...))
}

func (self androidLogger) Fatal(tag string, message string) {
	if self.enabled && self.severity <= C.ANDROID_LOG_FATAL {
		self.androidLog(C.ANDROID_LOG_FATAL, tag, message)
	}
}

func (self androidLogger) Fatalf(tag string, format string, args ...interface{}) {
	self.Fatal(tag, fmt.Sprintf(format, args...))
}

func (self *androidLogger) Write(p []byte) (int, error) {
	n := len(p)
	limit := bytes.IndexByte(p, '\n')
	for limit >= 0 {
		self.writeOneLine(p[:limit])
		p = p[limit+1:]
		limit = bytes.IndexByte(p, '\n')
	}
	self.writeOneLine(p)
	return n, nil
}
func (self *androidLogger) writeOneLine(p []byte) {
	if self.enabled {
		self.androidLog(C.ANDROID_LOG_VERBOSE, "io.Writer(Android)", string(p))
	}
}

func cstring(s string) *C.char {
	b, err := unix.BytePtrFromString(s)
	if err != nil {
		b := [1]C.char{}
		return &b[0]
	}
	return (*C.char)(unsafe.Pointer(b))
}

func newNativeLogger(enabled bool, severity Severity) NativeLogger {
	switch severity {
	case VERBOSE:
		severity = C.ANDROID_LOG_VERBOSE
	case DEBUG:
		severity = C.ANDROID_LOG_DEBUG
	case INFO:
		severity = C.ANDROID_LOG_INFO
	case WARN:
		severity = C.ANDROID_LOG_WARN
	case ERROR:
		severity = C.ANDROID_LOG_ERROR
	case FATAL:
		severity = C.ANDROID_LOG_FATAL
	}

	return &androidLogger{
		enabled:  enabled,
		severity: severity,
	}
}
