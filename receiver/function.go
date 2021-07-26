// summary : 関数はそのままの関数

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// define receiver type
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
