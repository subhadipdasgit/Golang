package main

import "fmt"

func operation(a, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}
func main() {
	var a, b int
	fmt.Print("Enter your numbers:=")
	fmt.Scan(&a, &b)
	sum, sub := operation(a, b)
	fmt.Println("Summation=", sum, "Subtraction =", sub)
}
