package main

/*
#cgo CFLAGS: -I ../../so
#cgo LDFLAGS: -L ../../so -lgo_test

#include <stdint.h>
#include "go_test.h"

int callOnMeGo_cgo(int in); // Forward declaration.
*/
import "C"

import (
	"fmt"
	"unsafe"
)

//export callOnMeGo
func callOnMeGo(in int) int {
	fmt.Println(in)
	return 0
}

func main() {
	C.go_runme((C.PCALLBACK_PROC)(unsafe.Pointer(C.callOnMeGo_cgo)), 54)
}
