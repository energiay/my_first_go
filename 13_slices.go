package main

import "fmt"

func main() {
	aSlice := []int{-1, 3, 4, 1, 7}
	fmt.Print("aSlice: ")
	printSlice(aSlice)

	fmt.Println("len: ", len(aSlice))
	fmt.Println("cap: ", cap(aSlice))

	reSlice := aSlice[1:3]
	fmt.Print("reSlice: ")
	printSlice(reSlice)

	reSlice[0] = -100
	reSlice[1] = -200

	fmt.Println("len: ", len(reSlice))
	fmt.Println("cap: ", cap(reSlice))

	printSlice(aSlice)
	printSlice(reSlice)

}

func printSlice(x []int) {
	for _, elem := range x {
		fmt.Print(elem, " ")
	}
	fmt.Println()
}