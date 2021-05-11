package main

import "fmt"

func main() {
	var age int
	var name string
	fmt.Print("Enter your First Name = ")
	fmt.Scan(&name)
	fmt.Print("Enter your Age = ")
	fmt.Scan(&age)
	if age > 18 {
		fmt.Print("Your are an adult")
	} else {
		fmt.Print("You are a teen now")
	}
}
