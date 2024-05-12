package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApp() {
	defer logging()
	fmt.Println("Run Application")
}

func main() {
	/**
	defer function adalah fungsi yang di eksekusi ketika suatu fungsi selesai di eksekusi
	*/

	runApp()
}