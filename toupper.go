package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		if r >= 97 && r <= 122 {
			runes[i] = r - 32
		}
	}
	return string(runes)
}
