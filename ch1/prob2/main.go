package prob2

import "sort"

func ArePermutations(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	as := []rune(a)
	bs := []rune(b)

	sort.Slice(as, func(i, j int) bool {
		return as[i] < as[j]
	})

	sort.Slice(bs, func(i, j int) bool {
		return bs[i] < bs[j]
	})

	return string(as) == string(bs)

}
