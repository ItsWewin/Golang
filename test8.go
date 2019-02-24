package main

import (
	"fmt"
)

type recttangle struct {
	x0, y0, x1, y1 int
	fill color.RGBA
}

func resizeReact(rect *recttangle, width, height int) {
	(*rect).x1 += width
	react.y1 += height
}

func  main()  {
	rect := recttangle{4, 8, 20, 10, color.RGBA{0xFF, 0, 0, 0xFF}}
	fmt.Println(react)
	resizeReact(&rect, 5, 5)
	fmt.Println(rect)
}