package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	registerUser("Budi", func (name string) bool {
		return name == "anjing"
	})

	blacklistUser := func (name string) bool  {
		return name == "anjing"
	}
	registerUser("anjing", blacklistUser)
}