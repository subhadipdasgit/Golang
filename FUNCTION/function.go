package main

import "fmt"

func checkValue(n1, n2 int) {
	if n1 > n2 {
		fmt.Print(n1, " is greater than ", n2)
	} else if n1 < n2 {
		fmt.Print(n1, " is smaller than ", n2)
	} else {
		fmt.Println(n1, " is equal to ", n2)
	}
}
func main() {
	var num1, num2 int
	fmt.Print("ENTER NUM 1 := ")
	fmt.Scanln(&num1)
	fmt.Print("ENTER NUM 2 := ")
	fmt.Scanln(&num2)
	checkValue(num1, num2)
}
