package main

import (
	"errors"
	"fmt"
)

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func getFirstDuplicate(elements []int) int {
	// Use map to record duplicates since look up is important
	encountered := map[int]bool{}

	for v := range elements {
		if encountered[elements[v]] == true {
			return elements[v]
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
		}
	}
	panicOnError(errors.New("No duplicate"))
	return -1
}

func main() {
	elements := []int{1, 2, 3, 1, 2, 4, 0}
	fmt.Println("your array: ", elements)

	// Test our method.
	result := getFirstDuplicate(elements)
	fmt.Println("first duplicate: ", result)
}
