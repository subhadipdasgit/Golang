package main

import "fmt"

func main() {
	var input int
	fmt.Print("Enter your number:=")
	fmt.Scanln(&input)
	for i := 0; i < input; i++ {
		fmt.Println("Loop", i)
	}
}
