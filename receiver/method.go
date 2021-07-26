// summary : メソッドはレシーバを持った関数のこと

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// set receiver pointer type
// ref. https://go-tour-jp.appspot.com/methods/4
// レシーバ自身を更新することが多いため、変数レシーバよりもポインタレシーバの方が一般的です。
// 値レシーバだと参照を渡せないので引数に対して変更をかけても、レシーバ自身を更新できない
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main () {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
