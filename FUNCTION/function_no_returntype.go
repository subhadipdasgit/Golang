package main

import "fmt"

func sqr(x int) {
	result := x * x
	fmt.Print("Square := ", result)
}

func main() {
	var x int
	fmt.Print("Enter your number :=")
	fmt.Scan(&x)
	sqr(x)
}
