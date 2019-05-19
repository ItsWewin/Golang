package main

import (
	"fmt"

	"./model"
)

func main() {
	student := model.NewStudent("tom", 99.8)
	fmt.Println(student)
	fmt.Println(student.GetScore())
}
