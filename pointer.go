package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	/**
	pointer adalah kemampuan membuat reference ke lokasi data yang sama
	tanpa menduplikasi data yang ada.
	*/
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1

	address2.City = "Bandung"
	fmt.Println(address2)
}