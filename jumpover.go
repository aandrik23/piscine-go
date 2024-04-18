package piscine

func JumpOver(str string) string {
	if str == "" {
		return "\n"
	}

	var result string
	for i, char := range str {
		if (i+1)%3 == 0 {
			result += string(char)
		}
	}
	if result == "" {
		return "\n"
	}
	return result + "\n"
}
