package main

import "fmt"

func main() {
	passGrade := 80

	if 80 > passGrade {
		fmt.Println("Lulus")
	} else {
		fmt.Println("Gagal")
	}

	// if short expression

	desc := "Learn"
	// length := len(desc)
	// if length > 5 {
	// 	// do something
	// }

	// is equals with

	if length := len(desc); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}