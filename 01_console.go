package main

import "fmt"

func main() {
    num := 5
    myFirstString := "my string"
    mySecondString := "my second string \n"

    fmt.Print(num, mySecondString, myFirstString)
    fmt.Println()
    fmt.Println(num, mySecondString, myFirstString)
    fmt.Println(num, " ", mySecondString, " ", myFirstString)
    fmt.Printf("%d %s %s", num, myFirstString, mySecondString)
}
