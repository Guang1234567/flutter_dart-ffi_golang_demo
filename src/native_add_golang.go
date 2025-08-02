package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
*/
import "C"
import (
	"time"
)

const (
	TAG = "native_add_golang.go"
)

// 全局变量声明
var logger NativeLogger

// 初始化（可选，根据AndroidLogger是否需要初始化参数）
func init() {
	logger = newNativeLogger(true, 2)
}

//export TestPlus
func TestPlus(a, b int, c string) int {
	logger.Debugf(TAG, "a: %d, b: %d, c: %s", a, b, c)
	return a + b + 700
}

//export TestPlusLongRunning
func TestPlusLongRunning(a, b int, c string) int {
	time.Sleep(3 * time.Second)
	logger.Debugf(TAG, "a: %d, b: %d, c: %s", a, b, c)
	return a + b + 77880
}

//export TestDivideByZero
func TestDivideByZero() int {
	time.Sleep(5 * time.Second)
	return TestDivideByZero321()
}

func TestDivideByZero321() int {
	return TestDivideByZero123()
}

func TestDivideByZero123() int {
	logger.Debug(TAG, "Go: Preparing to divide by zero...")
	a := 10
	b := 0
	// 下面这行代码会引发一个运行时 panic：integer divide by zero
	c := a / b
	return c
}
