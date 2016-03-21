package main

import "fmt"

func main() {
	// Define a "string" here with your name here

	// Define an "int" as you age here.

	// Define a "bool" to tell if you're a student.

	// Replace these values with your variables
	fmt.Printf(PrintNameAndAge("Change me!", 0))
	fmt.Printf(PrintStudentStatus(true))
}

func PrintNameAndAge(name string, age int) string {
	// Return a formatted string with fmt.Sprintf!
	return fmt.Sprintf("Change this to print your name and age!")
}

func PrintStudentStatus(student bool) string {
	// Create an if statement to print if you're a student.
	return fmt.Sprintf("Are you a student?")
}
