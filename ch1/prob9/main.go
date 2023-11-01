package prob9

import "strings"

func IsRotation(s1, s2 string) bool {
	scratch := s1 + s1

	return strings.Contains(scratch, s2)
}
