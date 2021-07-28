package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// fmt.Stringer と同様に、 fmt パッケージは、変数を文字列で出力する際に error インタフェースを確認します。
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
