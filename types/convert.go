package main

import (
	"fmt"
	"math"
)

func main() {
	// var v, type T => T(v) change type of var v to type T
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
