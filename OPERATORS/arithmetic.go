package main

import "fmt"

func main() {
	a, b := 6, 4
	fmt.Println("addition =", a+b)
	fmt.Println("subtraction=", a-b)
	fmt.Println("multiplication=", a*b)
	fmt.Println("division=", a/b)
	fmt.Println("remainder", a%b)
	a += 1
	b -= 1
	fmt.Println("a+1=", a)
	fmt.Println("b-1=", b)
}
