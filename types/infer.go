package main

import "fmt"

func main() {
	v := 42 // go infer int type of this statement
	fmt.Printf("v is of type %T\n", v)
}
