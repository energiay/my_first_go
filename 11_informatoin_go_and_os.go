package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func getVersionGo() {
	myVersion := runtime.Version()

	fmt.Println("")
	fmt.Println(myVersion)

	major := strings.Split(myVersion, ".")[0][2]
	fmt.Println("major:", string(major))

	minor := strings.Split(myVersion, ".")[1]
	fmt.Println("minor:", minor)
	m1, _ := strconv.Atoi(string(major))
	m2, _ := strconv.Atoi(minor)

	if m1 == 1 && m2 < 19 {
		fmt.Println("Need Go version 1.19 or higher!")
	}

	fmt.Println("You are using Go version 1.19 or higher!")
}

func main() {
	fmt.Println("You are using", runtime.Compiler)
	fmt.Println("on", runtime.GOARCH, "machine")
	fmt.Println("Using Go version", runtime.Version())
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Number of Goroutines: ", runtime.NumGoroutine())

	getVersionGo()
}
