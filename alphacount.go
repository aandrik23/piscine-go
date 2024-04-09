package piscine

func AlphaCount(s string) int {
	number := 0
	for _, letters := range []rune(s) {
		if (letters >= 'A' && letters <= 'Z') || (letters >= 'a' && letters <= 'z') {
			number = number + 1
		}
	}
	return number
}
