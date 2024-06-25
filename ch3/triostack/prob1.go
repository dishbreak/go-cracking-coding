package triostack

import "errors"

var (
	d      [30]int
	stacks = [3][]int{
		d[0:10],
		d[10:20],
		d[20:30],
	}
	stackLen = [3]int{0, 0, 0}
	ptrs     = [3]int{
		0, 0, 0,
	}
)

func Empty(idx int) bool {
	return stackLen[idx] == 0
}

func Peek(idx int) int {
	if Empty(idx) {
		return 0
	}
	return stacks[idx][ptrs[idx]]
}

func Pop(idx int) int {
	if Empty(idx) {
		return 0
	}

	v := d[ptrs[idx]]
	stacks[idx][ptrs[idx]] = 0
	stackLen[idx]--
	if stackLen[idx] > 0 {
		ptrs[idx]--
	}
	return v
}

func Push(idx int, v int) error {
	if stackLen[idx] == 10 {
		return errors.New("stack full")
	}

	ptrs[idx]++
	stacks[idx][ptrs[idx]] = v
	stackLen[idx]++
	return nil
}
