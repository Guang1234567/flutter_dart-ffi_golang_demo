package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
*/
import "C"
import "fmt"

//export TestPlus
func TestPlus(a, b int, c string) int {
    fmt.Printf("a: %d, b: %d, c: %s\n", a, b, c)
    return a + b
}