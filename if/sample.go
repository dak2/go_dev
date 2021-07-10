package main

import (
	"fmt"
)

func divide(num int) string {
	res := num % 2
	if res == 1 {
		return "odd"
	} else if res == 0 {
		return "even"
	} else {
		return "other"
	}
};

func main() {
	fmt.Println(divide(3))
}
