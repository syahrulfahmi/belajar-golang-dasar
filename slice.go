package main

import "fmt"

func main() {
	/**
	- slice adalah potongan dari array
	- slice mirip dengan array, yang membedakan adalah
	- ukuran pada slice bisa berubah
	Detail tipe data slice
		- Pointer, petunjuk data pertama
		- Length, panjang dari slice
		- Capacity, kapasitas dari array dimana length tidak boleh
		  lebih dari capacity.

	- array[low:high] , create slice from low to index before high
	- array[low:], create slice from low to last index
	- array[:high], create slice from 0 to index before high
	- array[:], create slice from 0 to last index in array
	*/

	brandCar := [...]string{"Honda", "BMW", "Toyota", "Mazda", "Hyundai", "Daihatsu"}
	sliceCar := brandCar[2:5]
	fmt.Println(sliceCar)

	slice1 := brandCar[:]
	fmt.Println(slice1)
	/**
	function slice
		- len(slice), get length of slice
		- cap(slice), get capacity of slice
		- append(slice, data) create new slice with add new data to last position,
		  if capacity is full then create new array.
		- make([]TypeData, length, capacity), create new slice
		- copy(destination, source), copy slice from source to destination
	*/

	lenbrandCar := len(sliceCar)
	fmt.Println(lenbrandCar)

	capBrandCar := cap(sliceCar)
	fmt.Println(capBrandCar)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Belajar"
	newSlice[1] = "Go-lang"
	fmt.Println(newSlice)
}