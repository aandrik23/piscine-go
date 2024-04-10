package piscine

func IsNumeric(s string) bool {
	for _, i := range []rune(s) {
		if i < '0' || i > '9' {
			return false
		}
	}
	return true
}
