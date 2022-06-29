package swissknife

// IntArrFind find element index in arr. -1 if not found
func IntArrFind(a []int64, x int64) int {
	for i, v := range a {
		if x == v {
			return i
		}
	}
	return -1
}

// IntArrContains indicates whether x is contained in a.
func IntArrContains(a []int64, x int64) bool {
	for _, v := range a {
		if x == v {
			return true
		}
	}
	return false
}
