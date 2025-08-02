//go:build darwin && ios
// +build darwin,ios

package main

/*
#cgo CFLAGS: -Werror -fmodules -fobjc-arc -x objective-c

@import Foundation;

static void nslog(char *str) {
	NSLog(@"%@", @(str));
}
*/
import "C"
import (
	"bytes"
	"fmt"
	"unsafe"

	"golang.org/x/sys/unix"
)

type iosLogger struct {
	enabled  bool
	severity Severity
}

func (self iosLogger) nslog(level string, tag string, message string) {
	C.nslog(fmt.Sprintf("%s %s: %s", level, tag, message))
}

func (self iosLogger) Verbose(tag string, message string) {
	if self.enabled && self.severity <= VERBOSE {
		self.nslog("V", tag, message)
	}
}

func (self iosLogger) Verbosef(tag string, format string, args ...interface{}) {
	self.Verbose(tag, fmt.Sprintf(format, args...))
}

func (self iosLogger) Debug(tag string, message string) {
	if self.enabled && self.severity <= DEBUG {
		self.nslog("D", tag, message)
	}
}

func (self iosLogger) Debugf(tag string, format string, args ...interface{}) {
	self.Debug(tag, fmt.Sprintf(format, args...))
}

func (self iosLogger) Info(tag string, message string) {
	if self.enabled && self.severity <= INFO {
		self.nslog("I", tag, message)
	}
}

func (self iosLogger) Infof(tag string, format string, args ...interface{}) {
	self.Info(tag, fmt.Sprintf(format, args...))
}

func (self iosLogger) Warn(tag string, message string) {
	if self.enabled && self.severity <= WARN {
		self.nslog("W", tag, message)
	}
}

func (self iosLogger) Warnf(tag string, format string, args ...interface{}) {
	self.Warn(tag, fmt.Sprintf(format, args...))
}

func (self iosLogger) Error(tag string, message string) {
	if self.enabled && self.severity <= ERROR {
		self.nslog("E", tag, message)
	}
}

func (self iosLogger) Errorf(tag string, format string, args ...interface{}) {
	self.Error(tag, fmt.Sprintf(format, args...))
}

func (self iosLogger) Fatal(tag string, message string) {
	if self.enabled && self.severity <= FATAL {
		self.nslog("F", tag, message)
	}
}

func (self iosLogger) Fatalf(tag string, format string, args ...interface{}) {
	self.Fatal(tag, fmt.Sprintf(format, args...))
}

func (self *iosLogger) Write(p []byte) (int, error) {
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
func (self *iosLogger) writeOneLine(p []byte) {
	if self.enabled {
		self.nslog("V", "io.Writer(iOS)", string(p))
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
	return &iosLogger{
		enabled:  enabled,
		severity: severity,
	}
}
