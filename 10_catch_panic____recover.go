package main

import (
	"fmt"
	"os"
)

func a() {
	fmt.Println("begin a()")
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Recovered inside a()")
		}
	}()
	fmt.Println("another log in a()")
	b()
	fmt.Println("end a()")

}

func b() {
	fmt.Println("begin b()")
	panic("Panic in b()!")
	fmt.Println("end b()")
}

func main() {
	fmt.Println("begin main()")

	if len(os.Args) == 1 {
		panic("Not enough arguments!")
	}

	a()
	fmt.Println("end main()")
}
