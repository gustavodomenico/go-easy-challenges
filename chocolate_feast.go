package main

func chocolateFeast(n int32, c int32, m int32) int32 {
	numberWrappers := n / c
	barsEaten := n / c
	for numberWrappers >= m {
		barsEaten = barsEaten + 1
		numberWrappers = numberWrappers - m
		numberWrappers = numberWrappers + 1
	}
	return barsEaten
}
