package main

import (
	"errors"
	"fmt"
	"time"

	"rsc.io/quote"

	"github.com/liuning0820/golang-basic/math"
	"github.com/liuning0820/golang-basic/datatype"

)



func e(v int) (int, error) {
	if v == 0 {
		return 0, errors.New("Zero cannot be used")
	} else {
		return 2 * v, nil
	}
}

func sum(a int32, b int32) (s int32) {
	s = a + b
	return s

}


// new type "Fruit"
// underlying type: struct
type Fruit struct {
	Name string
	Price float32
}

func main() {

	now := time.Now()
	fmt.Println(now.UTC())

	var x int32
	var y uint8

	x = 512

	y = uint8(x)

	println(x, y)

	fmt.Printf("hello, world\n")
	fmt.Println("hello, Gophers!")

	// log.Fatal("Exception occured!")
	//log.Panicln("A panic wiht a newline!")

	//var apple Fruit
	var apple = Fruit{Name: "Apple", Price: 12.5}
	fmt.Println(apple)
	fmt.Println(apple.Price)

	pear:= Fruit{
		Name: "Pear",
		Price: 8,
	}

	fmt.Println("The ",pear.Name, "price is ", pear.Price)

	var banana = new(Fruit)
	banana.Name = "Banana"
	fmt.Println(banana)


	v, err := e(0)

	if err != nil {

		fmt.Println(err, v) // Zero cannot be used 0

	}

	fmt.Println(quote.Go())

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

	var pointerVar = datatype.Pointer()
	fmt.Println(pointerVar)

	var arrayVar = datatype.Array()
	fmt.Println(arrayVar)

	datatype.Slices()

}
