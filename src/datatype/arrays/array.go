package main

import (
	"fmt"
	"reflect"
)

// Array - Play With Array
func Array() int {

	var arr [4]int
	arr[0] = 1
	println(arr[0])
	// Arrays do not need to be initialized explicitly; the zero value of an array is a ready-to-use array whose elements are themselves zeroed
	println(arr[1])

	// Get the type of the array
	arrType := reflect.TypeOf(arr)
	fmt.Printf("Type of arr: %s\n", arrType)
	return arr[1]

}

// Sum calculates the total from an array of numbers.
func Sum(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}

	// for _, number :=range numbers{
	// 	sum +=number
	// }

	return sum
}

func main() {

	Array()

	b := [2]string{"Penn", "Teller"}

	println(b[0])
}
