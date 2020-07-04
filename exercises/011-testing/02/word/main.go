package word

import "strings"

// UseCount make wordmap
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count function
func Count(s string) int {
	// write the code for this func
	xs := strings.Fields(s)

	return len(xs)
}
