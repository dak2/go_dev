package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// v is local variable in if scope
	// can use in if block or else block
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(pow(3, 2, 10))
}
