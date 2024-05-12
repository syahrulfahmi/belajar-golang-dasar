package main

import "fmt"

func main() {
	type noKtp string

	var ktpOrang noKtp = "1111111"
	contoh := "2222222"
	var contohKtp = noKtp(contoh)

	fmt.Println(ktpOrang)
	fmt.Println(contohKtp)

}