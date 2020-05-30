package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Go 没有类。不过你可以为结构体类型定义方法。
// 方法就是一类带特殊的 接收者 参数的函数。
// 方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。
func (v Vertex)  Abs2() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
	fmt.Println(v.Abs2())
}

// 方法即函数
// 方法只是个带接收者参数的函数。