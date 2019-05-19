package main

import "errors"

func main() {
	validateName("test2")
}

func validateName(name string) string {
	if name != "test" {
		err := errors.New("the value is not expected!")
		panic(err)
	} else {
		return name
	}
}
