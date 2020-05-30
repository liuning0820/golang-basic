package main

import (
	"fmt"

	"github.com/liuning0820/golang-basic/math"
	"github.com/liuning0820/golang-basic/primitivedatatype"
)

func main() {
	fmt.Printf("hello, world\n")
	fmt.Println("hello, Gophers!")

	// short variables declaration
	lang, year := "Go Language", "2017"
	println(lang, "is released at ", year)

	var sum int32 = sum(3, 4)
	fmt.Println(sum)

	println(15 == 0xF)

	// play with Pointer
	// declare and dereference

	// var firstName *string
	// firstName = "Arthur"
	// println(firstName)

	// var firstName *string
	// *firstName = "Arthur"
	// println(firstName)

	// var firstName *string = new(string)
	// *firstName = "Arthur"
	// println(firstName)

	var firstName *string = new(string)
	*firstName = "Arthur"
	println(*firstName)

	//address
	secondName := "Toby"
	fmt.Println(secondName)

	ptr := &secondName
	fmt.Println(ptr, *ptr)

	secondName = "Jerry"
	fmt.Println(ptr, *ptr)

	var mathCal = math.Add(1, 2)
	println(mathCal)

	a, b := math.Swap("hello", "world")
	fmt.Println(a, b)

	var pointerVar = primitivedatatype.Pointer()
	fmt.Println(pointerVar)

	var arrayVar = primitivedatatype.Array()
	fmt.Println(arrayVar)

	primitivedatatype.Slices()

}

// Function
func sum(a int32, b int32) (s int32) {
	s = a + b
	return s

}
