package main

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"rsc.io/quote"

	"github.com/liuning0820/golang-basic/src/datatype"
	"github.com/liuning0820/golang-basic/src/datatype/v1"
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
	Name  string
	Price float32
}

const (
	Bytes = 1
	KB    = 1024
	MB    = KB * 1024
	GB    = MB * 1024
	TB    = GB * 1024
)

var (
	units       = []int{Bytes, KB, MB, GB, TB}
	unitStrings = []string{"B", "K", "M", "G", "T"}
)

func getReduce(unit string, n int64) (float64, string) {
	reduce := 0
	switch unit {
	case "B":
		reduce = 0
	case "K":
		reduce = 1
	case "M":
		reduce = 2
	case "G":
		reduce = 3
	case "T":
		reduce = 4
	}
	for {
		if reduce <= 0 {
			break
		}

	}

	return float64(n) / float64(units[reduce]), unitStrings[reduce]
}

func main() {

	var result = strconv.FormatInt(13, 16)

	fmt.Println(result)

	val, reduceUnit := getReduce(unitStrings[1], 1000000)

	fmt.Printf("val %q, reduceUnit: %q", val, reduceUnit)

	fmt.Println(quote.Go())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		// simulate work
		time.Sleep(10 * time.Second)
		fmt.Println("Job1 completed")
		wg.Done()
	}()

	go func() {
		// simulate work
		time.Sleep(5 * time.Second)
		fmt.Println("Job2 completed")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Done")

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

	pear := Fruit{
		Name:  "Pear",
		Price: 8,
	}

	fmt.Println("The ", pear.Name, "price is ", pear.Price)

	var banana = new(Fruit)
	banana.Name = "Banana"
	fmt.Println(banana)

	v, err := e(0)

	if err != nil {

		fmt.Println(err, v) // Zero cannot be used 0

	}

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

	var pointerVar = datatypeV1.Array()
	fmt.Println(pointerVar)

	var arrayVar = datatypeV1.Array()
	fmt.Println(arrayVar)

	datatype.Slices()

}
