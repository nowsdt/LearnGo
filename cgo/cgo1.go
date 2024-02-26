package main

// #include <stdio.h>
// #include <stdlib.h>

import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(Random())
	Print("c fputs")
}

func Random() int {
	return int(C.random())
}

func Print(s string) {
	cs := C.CString(s)
	C.fputs(cs, (*C.FILE)(C.stdout))
	C.free(unsafe.Pointer(cs))
}
