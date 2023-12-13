package main

import (
	"fmt"

	"github.com/liuning0820/golang-basic/internal/math"
)

func main() {
	fmt.Printf("hello world\n")
	fmt.Print(hello("hello world"))

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// If you only want a subset of the returned values, use the blank identifier _.

	_, c := vals()
	fmt.Println(c)

	// 引用其他package
	fmt.Println(math.Add(1, 2))

}
