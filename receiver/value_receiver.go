package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// method with value receiver
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	// メソッドが値レシーバの場合、p.Abs() は (*p).Abs() として解釈されます。
	// &で変数のアドレスを取得しておいて、*でアドレスの値を取得してそれをレシーバに取る
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
