package prob1

import "sort"

func HasUniqueChars(s string) bool {
	chars := make(map[rune]bool)

	for _, c := range s {
		if chars[c] {
			return false
		}
		chars[c] = true
	}

	return true
}

func HasUniqueCharsNoDataStructure(s string) bool {
	if len(s) < 2 {
		return true
	}

	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	for i := 1; i < len(b); i++ {
		if b[i-1] == b[i] {
			return false
		}
	}

	return true
}
