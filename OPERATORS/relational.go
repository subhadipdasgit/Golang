package main

import "fmt"

func main() {
	a, b := 10, 5
	fmt.Println("a equals b=", a == b)
	fmt.Println("a not equals b=", a != b)
	fmt.Println("a greater than b=", a > b)
	fmt.Println("a smaller than b=", a < b)
	fmt.Println("a equal smaller than b=", a <= b)
	fmt.Println("a equal greater than b=", a >= b)
}
