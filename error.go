package main

import (
	"errors"
	"fmt"
)

func divide(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan Nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main () {
	result, err := divide(10,0)
	if err == nil {
		fmt.Println("Hasil: ", result)
	} else {
		fmt.Println("Error", err.Error())
	}
}