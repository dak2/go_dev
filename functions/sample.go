package main

import "fmt"

// type defined after variables
func add(x int, y int) int {
	return x + y
}

// emit type definitions when type is duplicate
func sub(x, y int) int {
	return x - y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(sub(42, 13))
}
