package main

import "fmt"

func main() {
	var arr = [6]int{6, 5, 8, 4, 2, 9}
	for i := 0; i < 6; i++ {
		fmt.Print(arr[i], " ")
	}
}
