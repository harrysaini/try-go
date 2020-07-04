package utils

// IsMapEqual check if maps are equal
func IsMapEqual(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	flag := true

	for k, v := range m1 {
		if v != m2[k] {
			flag = false
			break
		}
	}

	return flag
}
