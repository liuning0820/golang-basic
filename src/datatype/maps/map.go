package main

import (
	"errors"
	"fmt"
)

func Maps() map[string]int {
	var mapScores = map[string]int{"GO": 98, "Python": 80, "Java": 78}
	fmt.Println(mapScores)

	for key, value := range mapScores {
		fmt.Printf("%q is the key for the value %d\n", key, value)
	}

	return mapScores

}

func Search(dictionary map[string]int, word string) int {
	return dictionary[word]
}

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search2(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func main() {
	//Create a map
	subjectMarks := map[string]float32{"Golang": 85, "Java": 80, "Python": 81}
	fmt.Println(subjectMarks)
	var subjectMarks2 = map[string]float32{"Golang": 85, "Java": 80, "Python": 81}
	fmt.Println(subjectMarks2)

	//Access Values of a Map
	fmt.Println(subjectMarks["Golang"])

	//Change value of a map
	subjectMarks["Python"] = 90
	fmt.Println(subjectMarks["Python"])
	fmt.Println(subjectMarks)

	//Add Element of Go Map Element.
	//In Go, the order of elements in a map is not guaranteed.
	subjectMarks["C#"] = 95
	subjectMarks["Shell"] = 88
	fmt.Println("Updated Map:", subjectMarks)

	//Retrieve the value
	fmt.Println("Retrieve the value:", subjectMarks["Shell"])
	fmt.Println("Retrieve the value not existed:", subjectMarks["AI"])

	//Get Length of Map
	fmt.Println("len of the map:", len(subjectMarks))

	//Delete
	delete(subjectMarks, "Shell")
	fmt.Println(subjectMarks)

}
