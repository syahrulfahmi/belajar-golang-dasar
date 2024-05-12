package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func runApplication(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
}

func main() {
	/**
	panic func adalah function yang akan menghentikan program akan tetapi defer
	function akan tetap di eksekusi
	*/
	runApplication(true)
}