package main

import "fmt"

func faktorial(number int) int {
	if number == 1 {
		return number
	} else {
		return number * faktorial(number-1)
	}
}

func main() {
	result := faktorial(5)
	fmt.Println("Hasil faktorial", result)
}