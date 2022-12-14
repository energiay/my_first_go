package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	begin := time.Now() // измерение времени

	logs := []string{
		"[2022-08-01 14:28] тестовый лог для тестовой задачи",
		"[2022-0-8-27 14:28] тестовый лог для тестовой задачи",
		"[2022-08-02 14:28] тестовый лог для тестовой задачи",
		"[20-22-08-27 14:28] тестовый лог для тестовой задачи",
		"testing[2022-08-03 14:28] тестовый лог для тестовой задачи",
	}

	for _, logEntry := range logs {
		//r := regexp.MustCompile(`.*(\d\d\d\d-\d\d-\d\d).*`)
		r := regexp.MustCompile(`.*(\d\d\d\d-\d\d-\d\d).*`)
		if r.MatchString(logEntry) {
			match := r.FindStringSubmatch(logEntry)
			//fmt.Println(match)

			dt, err := time.Parse("2006-01-02", match[1])
			if err == nil {
				newFormat := dt.Format(time.RFC850)
				fmt.Println(newFormat)
			} else {
				fmt.Println("no:", err)
			}
		} else {
			//fmt.Println("Not a match!")
		}
	}

	duration := time.Since(begin)
	fmt.Println("It took this code: <", duration, "> to finish")
}
