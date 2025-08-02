package main

/*
#include <stdlib.h>

*/
import "C"
import (
	"io"
)

type Severity C.int

const (
	VERBOSE Severity = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

type NativeLogger interface {
	Verbose(tag string, message string)
	Verbosef(tag string, format string, args ...interface{})

	Debug(tag string, message string)
	Debugf(tag string, format string, args ...interface{})

	Info(tag string, message string)
	Infof(tag string, format string, args ...interface{})

	Warn(tag string, message string)
	Warnf(tag string, format string, args ...interface{})

	Error(tag string, message string)
	Errorf(tag string, format string, args ...interface{})

	Fatal(tag string, message string)
	Fatalf(tag string, format string, args ...interface{})

	io.Writer
}
