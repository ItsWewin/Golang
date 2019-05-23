package main

import "fmt"

func main() {
	const (
		a = 3
		b
		k = 12.5
		t = "abc"
		h
		c, d, z = iota, iota, iota
		e       = 12.5
		g
		x
		y
	)

	const (
		f = iota
	)

	fmt.Printf("a: %d, b: %d, c: %d, d: %d, e: %v\n", a, b, c, d, e)
	fmt.Printf("g: %v\n", g)
	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v\n", y)
	fmt.Printf("f: %d\n", f)
	fmt.Printf("z: %d\n", z)
	fmt.Printf("k: %d\n", k)
	fmt.Printf("h: %v\n", h)

}
