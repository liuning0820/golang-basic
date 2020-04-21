package main

import (
	"fmt"

	"./primitivedatatype"
)

func main() {
    fmt.Printf("hello, world\n")
    fmt.Println("hello, Gophers!")

    // short variables declaration
    lang, year := "Go Language", "2017"
    println(lang,"is released at ", year)

    var sum int32 = sum(3,4)
    fmt.Println(sum)

    println(15==0xF)

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
    secondName :="Toby"
    fmt.Println(secondName)

    ptr := &secondName
    fmt.Println(ptr, *ptr)

    secondName ="Jerry"
    fmt.Println(ptr, *ptr)


}

// Function
func sum(a int32, b int32)(s int32){
    s = a+b
    return s

}



