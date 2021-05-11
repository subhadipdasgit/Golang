package main

import (
	"fmt"
	"math"
)

func main() {
	var i, j int = 3, 4
	var k float64 = math.Sqrt(float64(i*i + j*j))
	var l uint = uint(k)
	fmt.Println(i, j, l)
}
