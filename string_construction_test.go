package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestStringConstructionCase1(t *testing.T) {
	var expected int32 = 4
	actual := stringConstruction("abcd")

	assert.Equal(t, actual, expected)
}
func TestStringConstructionCase2(t *testing.T) {
	var expected int32 = 2
	actual := stringConstruction("abab")

	assert.Equal(t, actual, expected)
}
