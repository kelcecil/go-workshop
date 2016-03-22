package main

import (
	"errors"
	"fmt"
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
	result := 0
	for i := range numbers {
		result += numbers[i]
	}
	return result
}

// Write a function that looks for a key in a map.
// Returns a value if available or an error if nil
func RetrieveNameAndNumberFromMap(name string) (string, error) {
	for k, v := range info {
		if k == name {
			return v, nil
		}
	}
	return "", errors.New(fmt.Sprintf("No key %v found", name))
}

// Write a function to append a number to an array.
func InsertNumberIntoArray(number int) {
	numbers = append(numbers, number)
}

// Write a function to add a value to a map.
func InsertNameAndPhoneIntoMap(name string, phone string) {
	info[name] = phone
}
