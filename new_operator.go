package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	/**
	- new adalah salah satu cara untuk membuat pointer
	- namun func new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal.
	*/
	alamat1 := new(Address)
	alamat2 := alamat1

	alamat2.Country = "Indonesia"
	fmt.Println(alamat1)
	fmt.Println(alamat2)
}