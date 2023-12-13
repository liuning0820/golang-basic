package datatype

import (
	"errors"
	"fmt"
)

func Maps() map[string](int) {
	var map_scores = map[string]int{"GO": 98, "Python": 80, "Java": 78}
	fmt.Println(map_scores)

	for key, value := range map_scores{
		fmt.Printf("%q is the key for the value %d\n", key, value)
	}

	return map_scores

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
