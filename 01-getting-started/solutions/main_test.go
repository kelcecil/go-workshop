package main

import "testing"

func TestPrintNameAndAge(t *testing.T) {
	sentence := PrintNameAndAge("Test", 30)
	if sentence == "Change this to print your name and age!" {
		t.Errorf("This method has yet to be implemented.")
	}
	if sentence == "My name is Kel, and I am 30 years old. I am a student! " {
		t.Fail()
	}
}

func TestStudentFunc(t *testing.T) {
	isStudent := PrintStudentStatus(true)
	isNotStudent := PrintStudentStatus(false)

	if isStudent == isNotStudent {
		t.Fatalf("This method has yet to be implemented")
	}
}
