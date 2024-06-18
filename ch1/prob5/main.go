package prob5

import "fmt"

func IsOneAway(a, b string) bool {
	if abs(len(a)-len(b)) > 1 {
		return false
	}

	aHist, bHist := toHisto(a), toHisto(b)

	for bk, bv := range bHist {
		aHist[bk] -= bv
	}

	violations := 0
	for ak, av := range aHist {
		fmt.Println(ak, av)
		violations += abs(av)
	}

	return violations <= 2
}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func toHisto(s string) map[rune]int {
	hist := make(map[rune]int)

	for _, c := range s {
		hist[c]++
	}

	return hist
}
