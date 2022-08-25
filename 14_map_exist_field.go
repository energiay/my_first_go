package main

import "fmt"

func main() {
	iMap := make(map[string]int)
	iMap["k1"] = 11
	iMap["k2"] = 22
	fmt.Println("iMap:", iMap)

	anotherMap := map[string]int{
		"k1": 11,
		"k2": 22,
	}
	fmt.Println("anotherMap:", anotherMap)

	_, ok := iMap["k1"]
	fmt.Println(ok)
	if ok {
		fmt.Println("exists")
	} else {
		fmt.Println("don't exists")
	}
}
