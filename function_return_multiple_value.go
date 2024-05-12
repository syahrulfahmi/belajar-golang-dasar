package main

import "fmt"

func getBrandWithVariant() (string, string) {
	return "Honda", "Brio"
}

func main() {
	brand, variant := getBrandWithVariant()
	
	/**
		skip "variant" return value
	*/
	// brand, _ := getBrandWithVariant()
	fmt.Println(brand, variant)
}