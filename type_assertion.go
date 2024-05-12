package main

import "fmt"

func random() interface{} {
	return false
}

func main() {
	/**
	type assertion adalah kemampuan merubah tipe menjadi tipe data yang diinginkan
	fitur ini seringkali digunakan ketika kita bertemu dengan data interface kosong
	*/
	// result := random().(string)
	// fmt.Println(result)

	// resultInt := random().(int)
	// fmt.Println(resultInt) // panic

	// cara yang lebih aman	
	val := random()
	switch value := val.(type) {
	case string :
		fmt.Println("String", value)
	case int:
		fmt.Println("integer", value)
	default:
		fmt.Println("Unknown", value)
	}

}