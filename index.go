package piscine

func Index(s string, toFind string) int {
	i := 0
	j := 0
	for range s {
		i++
	}
	for range toFind {
		j++
	}
	if j == 0 {
		return 0
	} else if j > i {
		return -1
	} else if j == i {
		if s == toFind {
			return 0
		}
		return -1
	} else if j < i {
		for h := 0; h <= i-j; h++ {
			if s[h:h+j] == toFind {
				return h
				break
			}
		}
	}

	return -1
}
