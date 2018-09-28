package diffcheck

// check the equivalent of each arrays
func arraysEqual(ar1, ar2 []string) bool {
	if len(ar1) != len(ar2) {
		return false
	}
	for i, c := range ar1 {
		if c != ar2[i] {
			return false
		}
	}
	return true
}
