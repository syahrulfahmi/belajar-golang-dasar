package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke ", counter)
	// 	counter++
	// }

	// maxCount := 10
	// for i := 0; i < maxCount; i++ {
	// 	fmt.Println("Perulangan ke ", i)
	// }

	brands := []string {"Honda", "Toyota", "BMW"}
	// for i := 0; i < len(brand); i++ {
	// 	fmt.Println(brands[i])
	// }

	// FOR RANGE
	for index, name := range brands {
		fmt.Println("index", index, "=", name)
	}

	for _, name := range brands {
		fmt.Println(name)
	}
}