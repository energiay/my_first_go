package main

// #include <stdio.h>
// void callC() {
// 	 printf("I am C code!\n");
// }
import "C"

import "fmt"

func main() {
	fmt.Println("I'm Go code")
	C.callC()
	fmt.Println("Another Go code")
}
