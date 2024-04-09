package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	for i, value := range runes {
		if i+1 == n {
			return value
		}
	}
	return 0
}
