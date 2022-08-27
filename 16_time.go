package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//fmt.Println("Epoch time:", time.Now().Unix())
	//
	//t := time.Now()
	//fmt.Println(t)
	//fmt.Println(t.Format(time.RFC3339))
	//fmt.Println(t.Weekday(), t.Day(), t.Month(), t.Year())

	//time.Sleep(time.Second)
	//t1 := time.Now()
	//fmt.Println("Time difference:", t1.Sub(t1))
	//
	//formatT := t.Format("01 January 2006")
	//fmt.Println(formatT)
	//
	//loc, _ := time.LoadLocation("Europe/Paris")
	//londonTime := t.In(loc)
	//fmt.Println("Paris:", londonTime)
	//
	var myTime string
	//if len(os.Args) != 2 {
	//	fmt.Println("usage : %s string", filepath.Base(os.Args[0]))
	//	os.Exit(1)
	//}

	myTime = os.Args[1]

	//d, err := time.Parse("15:04", myTime) // 15 - часы, 04 - минуты, 05  - секунды
	//if err == nil {
	//	fmt.Println("Full:", d)
	//	fmt.Println("Time:", d.Hour(), d.Minute())
	//} else {
	//	fmt.Println(err)
	//}

	d1, err1 := time.Parse("02 January 2006", myTime)
	if err1 == nil {
		fmt.Println("Full:", d1)
		fmt.Println("Time:", d1.Day(), d1.Month(), d1.Year())
	} else {
		fmt.Println(err1)
	}

}
