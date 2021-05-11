package main

import "fmt"

const MAX int = 3

func main() {
	a := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int
	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i]
	}
	for i = 0; i < MAX; i++ {
		fmt.Printf("Value of a [%d]= [%d] \t Address of [%d] = [%x]\n", i, *ptr[i], i, ptr[i])
	}
}
