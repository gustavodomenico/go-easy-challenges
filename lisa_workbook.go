package main

import "fmt"

// https://www.hackerrank.com/challenges/lisa-workbook/problem?isFullScreen=true

func workbook(n int32, k int32, arr []int32) int32 {
	result := int32(0)
	pages := int32(1)

	for i := int32(0); i < n; i++ {
		problems := arr[i]

		for j := int32(1); j <= problems; j++ {
			fmt.Printf("\n%d - chapter %d, problem %d, page %d", problems, i+1, j, pages)

			if j == pages {
				result = result + 1
				fmt.Printf("- SPECIAL -")
			}

			// End of the capacity
			if j != problems && j%k == 0 {
				pages = pages + 1
			}

		}

		pages = pages + 1
	}

	return result
}
