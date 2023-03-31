package main

import "strconv"

// https://www.hackerrank.com/challenges/fair-rations/problem?isFullScreen=true

func fairRations(B []int32) string {
	var breads int = 0
	for i := 0; i < len(B)-1; i++ {
		if B[i]%2 == 1 {
			B[i] = B[i] + 1
			B[i+1] = B[i+1] + 1

			breads = breads + 2
		}
	}

	if B[len(B)-1]%2 == 1 {
		return "NO"
	}

	return strconv.Itoa(breads)
}
