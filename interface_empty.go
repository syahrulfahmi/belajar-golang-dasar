package main

import "fmt"

func Ups() interface{} {
	return "Ups"
}

func main() {
	fmt.Println(Ups())
}