package main

import "fmt"

func getBrand(brand string) string {
	return "Ini Brand " + brand
}

func main() {
	result := getBrand("Honda")
	fmt.Println(result)

}