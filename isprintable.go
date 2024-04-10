package piscine

func IsPrintable(s string) bool {
	for _, i := range []rune(s) {
		if (i < 'A' || i > 'Z') && (i < 'a' || i > 'z') && (i == '\n') {
			return false
		}
	}
	return true
}
