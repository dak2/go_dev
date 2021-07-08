package main

import "fmt"

// return multiple values by conma
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a,b)
}
