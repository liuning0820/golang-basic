package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var agePaul, ageJohn int = rand.Intn(100), rand.Intn(100)
	fmt.Println("Paul is ", agePaul, "years old")
	fmt.Println("John is ", ageJohn, "years old")

	if agePaul > ageJohn {
		fmt.Println("Paul is older than John")
	} else {
		fmt.Println("Paul is yonger than John, or both have the same age")
	}

	if agePaul > ageJohn {
		fmt.Println("Paul is older than John")
	} else if agePaul == ageJohn {
		fmt.Println("Paul and John have the same age")

	} else {
		fmt.Println("Paul is yonger than John")
	}

	if agePaul > ageJohn {
		fmt.Println("Paul is older than John")
	} else {
		if agePaul == ageJohn {
			fmt.Println("Paul and John have the same age")
		} else {
			fmt.Println("Paul is yonger than John")
		}
	}
}
