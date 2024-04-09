package piscine

func IsUpper(s string) bool {
	for _, i := range []rune(s) {
		if i < 'A' || i > 'Z' {
			return false
		}
	}
	return true
}
