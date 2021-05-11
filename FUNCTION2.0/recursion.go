package main

import "fmt"

func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}
func main() {
	var i int
	fmt.Print("Enter number : ")
	fmt.Scan(&i)
	fmt.Println(" factorial of", i, "is", factorial(i))
}
