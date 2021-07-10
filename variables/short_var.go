package main

import "fmt"

func main() {
	var i, j int = 1, 2
	// can init variable without var statement in functions
	// out of functions, need var statement
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
