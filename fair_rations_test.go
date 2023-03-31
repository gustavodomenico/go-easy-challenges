package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestFairRationsCase1(t *testing.T) {
	expected := "4"
	actual := fairRations([]int32{4, 5, 6, 7})

	assert.Equal(t, actual, expected)
}
func TestFairRationsCase2(t *testing.T) {
	expected := "4"
	actual := fairRations([]int32{2, 3, 4, 5, 6})

	assert.Equal(t, actual, expected)
}

func TestFairRationsCase3(t *testing.T) {
	expected := "NO"
	actual := fairRations([]int32{1, 2})

	assert.Equal(t, actual, expected)
}
