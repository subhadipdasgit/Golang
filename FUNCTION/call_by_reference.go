package main

import "fmt"

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp

}
func main() {
	var a, b int
	fmt.Print("enter the numbers =  ")
	fmt.Scan(&a, &b)
	fmt.Println("Before Swap:=  ", a, b)
	swap(&a, &b)
	fmt.Println("after Swap:=  ", a, b)
}
