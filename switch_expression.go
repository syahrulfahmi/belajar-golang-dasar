package main

import "fmt"

func main() {
	brand := "Honda"

	switch brand {
	case "Honda":
		fmt.Println("Hello Honda")
	case "Toyota" :
		fmt.Println("Hello Toyota")
	default:
		fmt.Println("Hello Semua")
	}

	// switch with short statement
	switch length := len(brand); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}
 }