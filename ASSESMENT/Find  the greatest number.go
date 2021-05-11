package main

import "fmt"

func main() {
	fmt.Println("FIND THE GREATEST NUMBER :")
	var first int
	fmt.Print("Enter the first number : ")
	fmt.Scanln(&first)
	var second int
	fmt.Print("Enter the second number : ")
	fmt.Scanln(&second)
	var third int
	fmt.Print("Enter the third number : ")
	fmt.Scanln(&third)
	if first >= second && first >= third {
		fmt.Print(first, " is the greatest among the three numbers")
	} else if second >= first && second >= third {
		fmt.Print(second, " is the greatest among the three numbers")
	} else {
		fmt.Print(third, " is the greatest number")
	}

}
