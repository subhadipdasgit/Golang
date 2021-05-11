package main

import "fmt"

func main() {
	a := 1
	var c int = a
	fmt.Println(c)
	c += a
	fmt.Println(c)
	c -= a
	fmt.Println(c)
	c *= a
	fmt.Println(c)
	c /= a
	fmt.Println(c)
	c %= a
}
