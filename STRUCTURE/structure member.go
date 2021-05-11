package main

import "fmt"

type Person struct {
	name  string
	age   int
	mail  string
	phone int
}

func main() {
	p := Person{"Alex", 23, "alex@test.com", 8745213698}
	fmt.Println("Name : ", p.name)
	fmt.Println("Age : ", p.age)
	fmt.Println("Mail : ", p.mail)
	fmt.Println("Phone : ", p.phone)
}
