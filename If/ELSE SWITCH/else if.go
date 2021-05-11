package main

import "fmt"

func main() {
	var age int
	fmt.Print("Enter your Age = ")
	fmt.Scan(&age)
	if age >= 18 && age < 60 {
		fmt.Print("You are an adult")
	} else if age >= 60 {
		fmt.Print("You are a Senior citizen")
	} else {
		fmt.Print("You are a teen")
	}

}
