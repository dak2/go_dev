package main

import "fmt"

// without explicitly writing return variables, when return value named
// x, y is defined as return value
// but I think writing explicity return variables
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
