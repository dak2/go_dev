package main

import "fmt"

func main() {
	i, j := 42, 2701

	// & operator is getter of pointer of variable
	// * operater is getter of value linked pointer
	p := &i         // point to i
	fmt.Println(p)  // output pointer
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i => i = 21

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
