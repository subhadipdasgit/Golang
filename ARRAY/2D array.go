package main

import "fmt"

func main() {
	var a = [3][2]int{{0, 1}, {2, 3}, {3, 4}}
	var i, j int
	for i = 0; i < 3; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}
