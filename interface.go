package main

import "fmt"

type HasBrand interface {
	GetName() string
}

func sayHello(hasBrand HasBrand) {
	fmt.Println("Hello", hasBrand.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	/**
	- interface adalah tipe data abstrack
	- interface biasanya digunakan sebagai kontrak, artinya harus ada yang mengimplementasikan.
	IMPLEMENTASI INTERFACE
	- setiap data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface
	- sehingga kita tidak perlu mengimplementasikan interface secara eksplisit
	*/
	person := Person{"Budi"}
	sayHello(person)
}