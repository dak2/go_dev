package main

import "fmt"

func main() {
	// return 0 without init values of int type
	// return false without init values of bool type
	// return "" without init values of string type
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
