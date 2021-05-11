package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("current time : ", t.Hour(), ":", t.Minute(), ":", t.Second())
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good evening")
	}

}
