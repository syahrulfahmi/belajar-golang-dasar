package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (v *notFoundError) Error() string {
	return v.Message
}

func saveData(id string, data interface{}) error {
	if id == "" {
		return &validationError{Message: "validation error"}
	}

	if data != "budi" {
		return &notFoundError{Message: "not found error"}
	}

	// ok
	return nil
}

func main() {
	err := saveData("budi", nil)
	if err != nil {
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println("Validation Err:", validationErr.Error())
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println("not found error: ", notFoundErr.Error())
		} else {
			fmt.Println("Unknown Error: ", err.Error())
		}
	} else {
		fmt.Println("Sukses")
	}
}