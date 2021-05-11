package main

import "fmt"

func main() {
	var name, surname string
	var age int
	fmt.Print("Enter your name = ")
	fmt.Scan(&name, &surname)
	fmt.Print("Enter your age = ")
	fmt.Scan(&age)
	fmt.Println("Your Name is", name, surname)
	fmt.Println("Your Age is ", age)
}
