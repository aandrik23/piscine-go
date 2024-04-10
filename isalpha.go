package piscine

func IsAlpha(s string) bool {
	for _, i := range []rune(s) {
		if (i < 'A' || i > 'Z') && (i < 'a' || i > 'z') && (i < '0' || i > '9') {
			return false
		}
	}
	return true
}
