package main

import "fmt"

func main() {
	var brandCar [3]string
	brandCar[0] = "Honda"
	brandCar[1] = "BWM"
	brandCar[2] = "Toyota"

	fmt.Println(brandCar[0])
	fmt.Println(brandCar[1])
	fmt.Println(brandCar[2])

	// create direct array
	var values = [3] int {
		80, 90, 82,
	}
	fmt.Println(values)

	// function array
	// len(array) get size of an array
	// array[i] get data from index
	// array[i] = value set data base on index

	var dynamicArr = [...] int {
		80, 90, 82,12,
	}
	sizeValues := len(dynamicArr)
	fmt.Println(sizeValues)
}