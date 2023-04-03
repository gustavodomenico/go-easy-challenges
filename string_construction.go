package main

// https://www.hackerrank.com/challenges/string-construction/problem?isFullScreen=true

func stringConstruction(s string) int32 {
	var set map[rune]rune = make(map[rune]rune)

	chars := []rune(s)
	for i := 0; i < len(chars); i++ {
		set[chars[i]] = chars[i]
	}

	return int32(len(set))
}
