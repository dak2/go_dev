package main

import "fmt"

type Vertex struct {
	X, Y float64
}

// method with pointer receiver
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// function
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	// メソッドがポインタレシーバの場合、Goは利便性のため、v.Scale(2) のステートメントを (&v).Scale(2)として解釈
	v.Scale(2)
	// ScaleFuncにポインタを渡さないとエラーになる
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
