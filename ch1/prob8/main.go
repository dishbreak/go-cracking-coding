package main

func ZeroOut(m [][]int) [][]int {
	rows, cols := make([]bool, len(m)), make([]bool, len(m[0]))

	for y, row := range m {
		for x, cell := range row {
			if cell != 0 {
				continue
			}

			rows[y] = true
			cols[x] = true
		}
	}

	for y, row := range m {
		for x := range row {
			if rows[y] || cols[x] {
				m[y][x] = 0
			}
		}
	}

	return m
}
