package main

import "fmt"

type Car struct {
	Brand, Variant string
	year           int
}

// struct method
func (car Car) getCarBrand() {
	fmt.Println("Ini adalah brand mobil: ", car.Brand)
}

func main() {
	/**
	- struct adalah template data yang digunakan untuk menggabungkan data dalam satu kesatuan
	- struct merupakan representasi data dalam program aplikasi
	- simple nya struct adalah kumpulan dari field
	- struct tidak bisa langsung digunakan
	- kita bisa membuat data/object dari struct yang kita buat
	- ada banyak cara membuat struct, salah satunya menggunakan struct literals
	*/

	var brio Car
	brio.Brand = "Honda"
	brio.Variant = "Brio"
	brio.year = 2023
	fmt.Println("Ini adalah mobil", brio.Variant)

	rush := Car {
		Brand: "Toyota",
		Variant: "Rush" ,
		year: 2023,
	}
	fmt.Println("Ini adalah mobil", rush.Variant)

	ayla := Car {"Daihatsu", "Ayla", 2023}
	fmt.Println("Ini adalah mobil", ayla.Variant)

	ayla.getCarBrand()
}