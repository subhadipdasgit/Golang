package main

import "fmt"

func main() {
	var cp, sp, amount float64
	fmt.Print("Enter the Cost Price :=  ")
	fmt.Scan(&cp)
	fmt.Print("Enter the Sale Price :=  ")
	fmt.Scan(&sp)
	if sp > cp {
		amount = ((sp - cp) / cp) * 100
		fmt.Println("Profit in percantage := ", amount, "%")
	} else if cp > sp {
		amount = ((cp - sp) / cp) * 100
		fmt.Println("Loss in percantage := ", amount, "%")

	} else {
		fmt.Println("No Profit or loss")
	}

}
