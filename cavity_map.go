package main

/* https://www.hackerrank.com/challenges/cavity-map/problem?isFullScreen=true */

func cavityMap(grid []string) []string {
	size := len(grid)
	result := grid
	for i, v := range grid {
		if i > 0 && i < size-1 {
			previousLine := grid[i-1]
			nextLine := grid[i+1]
			for j, c := range v {
				if j > 0 && j < size-1 {
					cell := int(c)
					if cell > int(previousLine[j]) && cell > int(nextLine[j]) && cell > int(v[j-1]) && cell > int(v[j+1]) {
						result[i] = result[i][:j] + "X" + result[i][j+1:]
					}
				}
			}
		}
	}
	return result
}
