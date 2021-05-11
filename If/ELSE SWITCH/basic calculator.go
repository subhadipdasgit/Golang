package main

import (
	"fmt"
	"os"
)

func main() {
	var num, num1 int
	var operator string
calculator:
	fmt.Print("Enter the no. = ")
	fmt.Scan(&num)
	fmt.Print("Enter the operator = ")
	fmt.Scan(&operator)
	fmt.Print("Enter the no. = ")
	fmt.Scan(&num1)
	switch operator {
	case "+":
		fmt.Println(num + num1)
	case "-":
		fmt.Println(num - num1)
	case "*":
		fmt.Println(num * num1)
	case "/":
		fmt.Println(num / num1)
	default:
		fmt.Println("Invalid operation")
	}

	var choice string
	fmt.Print("Enter your choice : = ")
	fmt.Scan(&choice)
	if choice == "no" {
		os.Exit(0)
	} else {
		goto calculator
	}

}
