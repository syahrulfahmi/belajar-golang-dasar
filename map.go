package main

import "fmt"

func main() {
	car := map[string]string{
		"brand": "Honda",
		"type":  "Brio",
	}

	fmt.Println(car["brand"])
	fmt.Println(car)
	fmt.Println(car["salah"]) // show blank output

	/**
	function map
		- len(map), get map size
		- map[key], get data with map key
		- map[key] = value, change data with map key
		- make(map[TypeKey]TypeValue), create new map
		- delete(map, key), delete map with key
	*/
	
	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["author"] = "James"
	fmt.Println(book)
}