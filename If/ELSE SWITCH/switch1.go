package main

import (
	"fmt"
	"time"
)

func main() {
	var n string
	fmt.Println("When is her birthday ?")
	fmt.Scan(&n)
	today := time.Now().Weekday() // time.Now() is used to convert time to weekday
	fmt.Println("Today is ", today)
	switch time.Friday {
	case today + 0:
		fmt.Println("Today is Friday")
	case today + 1:
		fmt.Println("Tommorrow")
	case today + 2:
		fmt.Println("Day after tommorow")
	case today + 3:
		fmt.Println("After 2 days")
	case today + 4:
		fmt.Println("After 3 Days is friday")
	default:
		fmt.Println("too far ")
	}
}
