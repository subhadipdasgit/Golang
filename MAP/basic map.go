package main

import "fmt"

func main() {
	var products map[string]int
	products = make(map[string]int)

	products["A"] = 20
	products["B"] = 50
	products["C"] = 70
	fmt.Println("Product of A value= ", products["A"])
	for prod := range products {
		fmt.Println(" Product of Index =", prod, "Product value =", products[prod])
	}
}
