package main

import (
	"fmt"
)

type Product struct {
	name string
	price float64
}

func (product Product) String() string {
	return fmt.Sprintf("%s (%.2f)", product.name, product.price)
}

func main() {
	products := []*Product{{"first", 3.99}, {"second", 1.99}, {"third", 0.99}}
	fmt.Println(products)

	for _, product := range products {
		product.price += 0.5
	}

	fmt.Print(products)
}