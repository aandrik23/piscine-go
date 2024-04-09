package piscine

func LastRune(s string) rune {
	last := []rune(s)
	lens := 0
	for range last {
		lens++
	}
	// 	return [lens (rune) -1]
	for i, value := range last {
		if i == lens-1 {
			return value
		}
	}
	return 0
}
