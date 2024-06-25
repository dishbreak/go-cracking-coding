package prob6

import (
	"fmt"
	"strings"
)

func Compress(s string) string {
	if len(s) < 3 {
		return s
	}

	s = s + " " //silly trick to ensure last character sequence gets written

	var sb strings.Builder

	ct := 1
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			ct++
			continue
		}
		sb.WriteByte(s[i])
		sb.WriteString(fmt.Sprint(ct))
		ct = 1
	}

	s = s[:len(s)-1]

	if sb.Len() >= len(s)-1 {
		return s
	}

	return sb.String()
}
