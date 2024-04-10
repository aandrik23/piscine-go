package piscine

func ToLower(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		if r >= 65 && r <= 90 {
			runes[i] = r + 32
		}
	}
	return string(runes)
}
