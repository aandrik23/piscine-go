package piscine

func f(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 2 {
		return true
	}
	sortedAsc := false
	sortedDesc := false
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) > 0 {
			sortedAsc = true
			if sortedDesc {
				return false
			}
		}
		if f(a[i-1], a[i]) < 0 {
			sortedDesc = true
			if sortedAsc {
				return false
			}
		}
		if f(a[i-1], a[i]) == 0 {
			sortedDesc = true
			sortedAsc = true
		} //
	}
	return sortedAsc || sortedDesc
}
