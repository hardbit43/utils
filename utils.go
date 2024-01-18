package utils

func InSlice(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func InSliseInt(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func InSliceFloat(a []float64, x float64) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
