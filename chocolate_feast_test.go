package main

import "testing"

func TestChocolateFeastCase1(t *testing.T) {
	expected := 9
	if actual := chocolateFeast(15, 3, 2); actual != int32(expected) {
		t.Errorf("The main function provided an actual of %d but %d was expected.", actual, expected)
	}
}
func TestChocolateFeastCase2(t *testing.T) {
	expected := 6
	if actual := chocolateFeast(10, 2, 5); actual != int32(expected) {
		t.Errorf("The main function provided an actual of %d but %d was expected.", actual, expected)
	}
}
