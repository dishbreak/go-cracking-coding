package prob7

import "image"

/*
111111
122221
123321
123321
122221
111111

ring 1
*/

var (
	north = image.Pt(0, -1)
	south = image.Pt(0, 1)
	east  = image.Pt(1, 0)
	west  = image.Pt(-1, 0)
)

func Rotate(m [][]int) [][]int {
	n := len(m)
	rings := n / 2

	for r := 0; r < rings; r++ {
		nw := image.Pt(r, r)
		ne := image.Pt(n-r, r)
		sw := image.Pt(r, n-r)
		se := image.Pt(n-r, n-r)
		for a, b, c, d := nw, ne, se, sw; a != ne; a, b, c, d = a.Add(east), b.Add(south), c.Add(west), d.Add(north) {
			m[a.Y][a.X], m[b.Y][b.X], m[c.Y][c.X], m[d.Y][d.X] = m[d.Y][d.X], m[a.Y][a.X], m[b.Y][b.X], m[c.Y][c.X]
		}
	}
	return m
}
