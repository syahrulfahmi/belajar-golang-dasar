package main

import "fmt"

/**
secara konsep mirip varargs di kotlin
*/

func sumAll(numbers ... int) int {
	total := 0
	for _, numbers := range numbers {
		total += numbers
	}
	return total
}

func main() {
	total := sumAll(1,2,3,4,5)
	fmt.Println(total)

	/*
	terkadang kita mengirimkan variadic args
	dalam tipe data slice
	*/
	numbers := []int {1,2,3,4,5}
	totalNumbers := sumAll(numbers...)
	fmt.Println(totalNumbers)
}