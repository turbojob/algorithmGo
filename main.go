package main

import (
	"fmt"
)

func main() {
	s1 := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(s1)

	fmt.Println(s1[0:len(s1)])
	fmt.Println(s1[0 : len(s1)-1])
}
