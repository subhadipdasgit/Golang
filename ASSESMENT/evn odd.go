package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter number : ")
	fmt.Scanln(&n)
	if n%2 == 0 {
		fmt.Println(n, "is Even number")
	} else {
		fmt.Println(n, "is Odd number")
	}
}
