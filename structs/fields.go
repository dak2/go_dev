package main

import "fmt"

type Square struct {
	width int
	height int
}

func main() {
	v := Square{1, 2}
	// can get width struct fields with dot(.)
	v.width = 4
	fmt.Println(v.width)
}
