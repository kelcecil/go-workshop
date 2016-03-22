package main

// Use "go test -v" to check your implementations!

import (
	"errors"
)

var (
	// Define a map with keys as strings and values as strings here
	info map[string]string = make(map[string]string)

	// Define an array of integers here
	numbers []int = make([]int, 0)
)

// Write a function using a for loop and the range function
// that adds the contents of an array and returns.
func SummationOfNumbersArray() int {
	return 0
}

// Write a function that looks for a key in a map.
// Returns a value if available or an error if nil
func RetrieveNameAndNumberFromMap(name string) (string, error) {
	return "", errors.New("Unimplemented!")
}

// Write a function to append a number to an array.
func InsertNumberIntoArray(number int) {

}

// Write a function to add a value to a map.
func InsertNameAndPhoneIntoMap(name string, phone string) {

}
