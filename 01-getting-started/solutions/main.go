package main

// Import the format package to get access to Sprintf and Printf below.
import "fmt"

func main() {
	// Define a "string" here with your name here
	// You can use the short-hand form to declare a variable!
	name := "Kel"

	// Define an "int" as you age here.
	// You can also use the var keyword.
	var age int = 30

	// Define a "bool" to tell if you're a student.
	student := true

	// Replace these values with your variables
	fmt.Printf(PrintNameAndAge(name, age))
	fmt.Printf(PrintStudentStatus(student))
}

func PrintNameAndAge(name string, age int) string {
	// Use fmt.Sprintf to easily format a format.
	// Use %v to use the value's default format.
	// Learn more: https://golang.org/pkg/fmt/

	return fmt.Sprintf("My name is %v, and I am %v years old. ", name, age)
}

func PrintStudentStatus(student bool) string {
	// Declare a string variable.
	var statement string

	// Set the value based on student.
	if student {
		statement = "I am a student!"
	} else {
		statement = "I am not a student!"
	}

	// Return the statement.
	return fmt.Sprintf(statement)
}
