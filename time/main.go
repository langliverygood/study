package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	fmt.Println(now.Format("1995-11-08"))

	timer := time.Tick(time.Second)
	for t := range timer {
		fmt.Println(t.Year())
	}

	timer2 := time.Tick(time.Second)
	for t := range timer2 {
		fmt.Println(t.Month())
	}
}
