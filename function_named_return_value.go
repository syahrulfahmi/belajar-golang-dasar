package main

import "fmt"

func getCompleteBrand() (brand, variant string, year int) {
	brand = "Honda"
	variant = "Brio"
	year = 2023
	return brand, variant, year
}

func main() {
	brand, variant, year := getCompleteBrand()
	fmt.Println(brand, variant, year)
}