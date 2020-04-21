package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
    fmt.Println("hello, Gophers!")

    // short variables declaration
    lang, year := "Go Language", "2017"
    println(lang,"is released at ", year)

    var sum int32 = sum(3,4)
    fmt.Println(sum)

    println(15==0xF)



}

// Function
func sum(a int32, b int32)(s int32){
    s = a+b
    return s

}



