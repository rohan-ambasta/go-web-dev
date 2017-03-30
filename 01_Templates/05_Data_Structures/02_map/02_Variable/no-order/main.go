package main

import (
	"fmt"
)

func main() {
	values := map[string]string{
		"captain":          "kohli",
		"vice-captain":     "rahane",
		"off-spinner":      "ashwin",
		"left-arm-spinner": "jadeja",
		"fast-bowler":      "umesh"}

	for key, val := range values {
		fmt.Println(key + " - " + val)
	}
}
