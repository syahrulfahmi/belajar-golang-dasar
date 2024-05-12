package main

import "fmt"

type Filter func(string) string // type declaration for shorter function name

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Budi", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}