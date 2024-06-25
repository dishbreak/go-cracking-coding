package prob4

import "strings"

func HasPalindromePermutation(s string) bool {
	if len(s) < 3 {
		return false
	}

	s = strings.ToLower(s)

	hits := make(map[rune]int)
	for _, c := range s {
		if c == ' ' {
			continue
		}
		hits[c]++
	}

	oddCt := 0
	for _, v := range hits {
		if v%2 == 1 {
			oddCt++
		}
	}

	return oddCt == 1 || oddCt == 0
}
