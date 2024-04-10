package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	for i, j := range s {
		if j >= 65 && j <= 90 {
			runes[i] = j + 32
		}
	}
	for r, n := range runes {
		if n >= 97 && n <= 122 && r != 0 {
			if (runes[r-1] < 65 || runes[r-1] > 90) && (runes[r-1] < 97 || runes[r-1] > 122) && (runes[r-1] < 0 || runes[r-1] > 9) {
				runes[r] = n - 32
			}
		} else if r == 0 && n >= 97 && n <= 122 {
			runes[r] = n - 32
		}
	}
	return string(runes)
}
