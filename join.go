package main

func join(strings []string, separator string) string {
	if len(strings) == 0 {
		return ""
	}
	s := ""
	lastIdx := len(strings) - 1
	for _, v := range strings[0:lastIdx] {
		s += v + separator
	}
	return s + strings[lastIdx]
}
