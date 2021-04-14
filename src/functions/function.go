package main

import (
	"fmt"
	"math"

	"github.com/liuning0820/golang-basic/src/constant"
)

type Vertex struct {
	x, y float64
}

func SqrtRoot(v Vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// Go 没有类。不过你可以为结构体类型定义方法。
// 方法就是一类带特殊的 接收者 参数的函数。
// 方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。
func (v Vertex) Abs2() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func CountryFromVin(vin string) string {
	switch vin[10] {
	case 'C':
		return constant.COUNTRY_CN
	case '3':
		return constant.COUNTRY_UK
	case 'R':
		return constant.COUNTRY_RESEARCH
	case '1', 'F', 'P':
		return constant.COUNTRY_US
	default:
		return constant.COUNTRY_UNKNOWN
	}
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(SqrtRoot(v))
	fmt.Println(v.Abs2())

	fmt.Println(CountryFromVin("5YJ3F7EA2KF461145"))
}

// 方法即函数
// 方法只是个带接收者参数的函数。
