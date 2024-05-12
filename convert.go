package main

import "fmt"

func main () {
	nilai32 := 32768
	nilai64 := int64(nilai32)
	nilai16 := int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	// convert string
	desc := "Belajar Go-Lang"
	e := desc[1]
	eString := string(e)

	fmt.Println(e)
	fmt.Println(eString)
}