package main

import "fmt"

func main() {
	var num int
	var language string
	fmt.Println("Welcome to Language selection")
	fmt.Print("Enter your choice : ")
	fmt.Scan(&num)
	switch num {
	case 1:
		language = "English"
	case 2:
		language = "Hindi"
	case 3:
		language = "Bengali"
	case 4:
		language = "Gujrathi"

	}
	fmt.Print("Language selected by you : ", language)
}
