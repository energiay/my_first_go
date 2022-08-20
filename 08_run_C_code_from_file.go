package main

// #cgo CFLAGS: -I${SRCDIR}/libs
// #cgo LDFLAGS: ${SRCDIR}/callC.a
// #include <stdlib.h>
// #include <callC.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("I'm Go code")

	C.cHello()

	message := C.CString("string")
	fmt.Println("message1: ", message)
	defer C.free(unsafe.Pointer(message)) // хз, зачем оно тут?
	fmt.Println("message2: ", message)
	C.printMessage(message)

	fmt.Println("Another Go code")
}

// ls -l libs
// gcc -c libs/*.c
// ls -l callC.o
// file callC.o
// /usr/bin/ar rs callC.a *.o
// ls -l callC.a
// file callC.a
