package main

import "fmt"

type myStrucrure struct {
	Name    string
	Surname string
	Height  int32
}

func createStruct(n, s string, h int32) *myStrucrure {
	if h > 300 {
		h = 0
	}

	return &myStrucrure{n, s, h}
}

func retStruct(n, s string, h int32) myStrucrure {
	if h > 300 {
		h = 0
	}

	return myStrucrure{n, s, h}
}

func main() {
	s1 := createStruct("Ivan", "Strausov", 123)
	s2 := retStruct("Ivan1", "Strausov1", 301)

	fmt.Println((*s1).Name)
	fmt.Println(s2.Name)
	fmt.Println(s1)
	fmt.Println(s2)

}
