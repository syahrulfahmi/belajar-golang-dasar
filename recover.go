package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	// contoh recover yang benar
	message := recover()
	fmt.Println("Terjadi panic", message)
}

func runApplication(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}

	// cara recover yang salah
	// message := recover()
	// fmt.Println("Terjadi Error", message)
}

func main() {
	/**
	recover adalah func yang bisa digunakan untuk menangkap data panic
	dengan recover proses panic akan terhenti, sehingga program akan tetap dijalankan
	*/
	runApplication(true)
}