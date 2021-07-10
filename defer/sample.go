package main

import "fmt"

func main() {
	// exec after fmt.Println("hello")
	defer fmt.Println("world")

	fmt.Println("hello")
}
