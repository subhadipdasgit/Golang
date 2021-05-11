package main

import "fmt"

func swap(x, y int) (int, int) {
	var temp int
	temp = x
	x = y
	y = temp
	return x, y
}
func main() {
	var a, b int
	fmt.Print("enter the numbers =  ")
	fmt.Scan(&a, &b)
	fmt.Println("Before Swap:=  ", a, b)
	a, b = swap(a, b)
	fmt.Println("after Swap:=  ", a, b)

}
