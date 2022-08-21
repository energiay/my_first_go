package main

import "fmt"

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
	a()
	fmt.Println("end main()")
}
