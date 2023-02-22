package main

import (
	"fmt"
)

func main() {

	test := make(map[int64]bool)
	test[0] = false
	test[1] = true
	fmt.Println(test[2])

	fmt.Println(map[int]bool{
		100: true,
		200: true,
	}[2])

}
