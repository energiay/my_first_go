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
}
