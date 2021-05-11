package main

import "fmt"

func sequencenumber() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
func main() {
	nextnumber := sequencenumber()
	fmt.Println(nextnumber())
	fmt.Println(nextnumber())
	nextnumber1 := sequencenumber()
	fmt.Println(nextnumber1())
	fmt.Println(nextnumber1())
	fmt.Println(nextnumber1())
}
