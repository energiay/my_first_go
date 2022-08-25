package main

import "fmt"

func main() {
	const (
		Zero int = iota
		One
		Two
		Three
		Four
	)
	fmt.Println(Zero)
	fmt.Println(One)
	fmt.Println(Four)

	const (
		p2_0 int = 1 << iota
		_
		p2_2
		_
		p2_4
		_
		p2_6
	)
	fmt.Println("2^0:", p2_0)
	fmt.Println("2^2:", p2_2)
	fmt.Println("2^4:", p2_4)
	fmt.Println("2^6:", p2_6)
}
