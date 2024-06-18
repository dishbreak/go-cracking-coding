package prob3

func URLify(s string) string {
	b := []byte(s)
	l := len(b)

	j := -1
	for i := l - 1; i >= 0; i-- {
		if b[i] != ' ' {
			j = i
			break
		}
	}

	for i := l - 1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if b[j] != ' ' {
			b[j], b[i] = b[i], b[j]
			continue
		}

		b[i] = '0'
		b[i-1] = '2'
		b[i-2] = '%'
		i -= 2
	}

	return string(b)
}
