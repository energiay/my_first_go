package main

import (
	"io"
	"os"
)

func main() {
	myStr := ""
	argument := os.Args
	if len(argument) == 1 {
		myStr = "Please give me one argument!"
	} else {
		myStr = argument[1] + " " + argument[0]
	}

	io.WriteString(os.Stdout, myStr)
	io.WriteString(os.Stdout, "\n")

	// вот это новость :)
	h, k := "4", "5"
	i, k := "2", "3"
	io.WriteString(os.Stdout, k)
	io.WriteString(os.Stdout, i)
	io.WriteString(os.Stdout, h)
	io.WriteString(os.Stdout, "\n")
}
