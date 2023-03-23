package main

import (
	"reflect"
	"testing"
)

func TestCavityMapCase1(t *testing.T) {
	expected := []string{"989", "1X1", "111"}
	if actual := cavityMap([]string{"989", "191", "111"}); !reflect.DeepEqual(actual, expected) {
		t.Errorf("The main function provided an actual of %s but %s was expected.", actual, expected)
	}
}
func TestCavityMapCase2(t *testing.T) {
	expected := []string{"1112", "1X12", "18X2", "1234"}
	if actual := cavityMap([]string{"1112", "1912", "1892", "1234"}); !reflect.DeepEqual(actual, expected) {
		t.Errorf("The main function provided an actual of %s but %s was expected.", actual, expected)
	}
}
