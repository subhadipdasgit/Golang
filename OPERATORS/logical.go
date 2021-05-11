package main

import "fmt"

func main() {
	a, b, c, d := 5, 6, 4, false

	fmt.Println((a > b) && (b > c))
	fmt.Println((a > b) || (b > c))
	fmt.Println(!d)

}
