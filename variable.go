package main

import "fmt"

func main() {
	var name string
	name = "Belajar Koding"
	fmt.Println(name)

	name = "Belajar Go-Lang"
	fmt.Println(name)

	// simplify
	desc := "Belajar Variabel"
	fmt.Println(desc)

	// multiple declare variable
	var (
		firstText = "Belajar"
		lastText = "Go-Lang"
	)
	fmt.Println(firstText)
	fmt.Println(lastText)

	// constant
	const address string = "Jakarta"
	const (
		firstDesc = "Belajar"
		lastDesc = "Go-Lang"
	)
	fmt.Println(address)
	fmt.Println(firstDesc)
	fmt.Println(lastDesc)
	
}