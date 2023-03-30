package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLisaWorkbookCase1(t *testing.T) {
	expected := int32(1)
	actual := workbook(2, 3, []int32{4, 2})

	assert.Equal(t, actual, expected)
}
func TestLisaWorkbookCase2(t *testing.T) {
	expected := int32(4)
	actual := workbook(5, 3, []int32{4, 2, 6, 1, 10})

	assert.Equal(t, actual, expected)
}

func TestLisaWorkbookCase3(t *testing.T) {
	expected := int32(8)
	actual := workbook(10, 5, []int32{3, 8, 15, 11, 14, 1, 9, 2, 24, 31})

	assert.Equal(t, actual, expected)
}
