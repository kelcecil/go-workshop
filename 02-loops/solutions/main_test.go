package main

import "testing"

func TestMapSize(t *testing.T) {
	InsertNamesIntoMap()
	if len(info) < 3 {
		t.Errorf("Need %v more items in map. Add more names and phone numbers!", (3 - len(info)))
	}
}

func TestMapRetrieval(t *testing.T) {
	InsertNamesIntoMap()
	phoneNumber, err := RetrieveNameAndNumberFromMap("Kel")
	if err != nil {
		t.Errorf("The key should exist.")
	}
	if phoneNumber != "304-867-5309" {
		t.Errorf("The value is not what was expected.")
	}
}

func TestSummation(t *testing.T) {
	InsertNumbersIntoArray()
	if SummationOfNumbersArray() != 6 {
		t.Errorf("Result does not equal six.")
	}
}

func InsertNamesIntoMap() {
	InsertNameAndPhoneIntoMap("Kel", "304-867-5309")
	InsertNameAndPhoneIntoMap("Test1", "304-111-1111")
	InsertNameAndPhoneIntoMap("Test2", "304-222-2222")
}

func InsertNumbersIntoArray() {
	InsertNumberIntoArray(1)
	InsertNumberIntoArray(2)
	InsertNumberIntoArray(3)
}
