package main

import (
	"fmt"
	"sort"
)

func main() {
	aSlice := []int{-1, 3, 4, 1, 7}
	fmt.Print("aSlice: ")
	printSlice(aSlice)

	fmt.Println("len: ", len(aSlice))
	fmt.Println("cap: ", cap(aSlice)) //capacity

	reSlice := aSlice[1:3]
	fmt.Print("reSlice: ")
	printSlice(reSlice)

	reSlice[0] = -100
	reSlice[1] = -200

	fmt.Println("len: ", len(reSlice))
	fmt.Println("cap: ", cap(reSlice))

	printSlice(aSlice)
	printSlice(reSlice)

	aSlice9 := []int{-1, 3, 4, 1, 7, 8, 5, 23, 4}
	array5 := [5]int{-9, 2, 3, 4, 5}
	copy(array5[0:], aSlice9)
	fmt.Println("array5:", array5)
	fmt.Println("aSlice9:", aSlice9)

	mySlice := make([]aStructure, 0)
	mySlice = append(mySlice, aStructure{"Mihalis", 180, 90})
	mySlice = append(mySlice, aStructure{"Mihalis", 160, 90})
	mySlice = append(mySlice, aStructure{"Mihalis", 380, 90})
	mySlice = append(mySlice, aStructure{"Mihalis", 185, 90})
	mySlice = append(mySlice, aStructure{"Mihalis", 10, 90})

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height > mySlice[j].height
	})
	fmt.Println(mySlice)
}

type aStructure struct {
	person string
	height int
	weight int
}

func printSlice(x []int) {
	for _, elem := range x {
		fmt.Print(elem, " ")
	}
	fmt.Println()
}
