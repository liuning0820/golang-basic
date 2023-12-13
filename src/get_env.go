package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Shell:", os.Getenv("SHELL"))
}
