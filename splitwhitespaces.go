package piscine

func SplitWhiteSpaces(str string) []string {
	str = word(str)
	len := wordlen(str)
	i := make([]string, len)
	k := ""
	j := 0

	if space(str[0]) {
		k += string(str[0])
	}

	for l := 1; l < strlen(str); l++ {
		if space(str[l-1]) {
			i[j] = k
			k = string(str[l])
			j++
		} else if space(str[l]) {
			k += string(str[l])
		}
	}
	i[j] = k

	return i
}

func word(s string) string {
	result := ""
	for l, v := range s {
		if l == 0 {
			if space(byte(v)) {
				result += string(v)
			}
		} else {
			if space(s[l-1]) && space(s[l]) {
				result += string(v)
			}
		}
	}
	len := strlen(result)
	if space(result[len-1]) {
		result = result[:len-1]
	}
	return result
}

func space(s byte) bool {
	return (s == ' ' || s == '\n' || s == '\t')
}

func wordlen(s string) int {
	if strlen(s) == 0 {
		return 0
	}
	count := 1
	for _, v := range s {
		if space(byte(v)) {
			count++
		}
	}
	return count
}

func strlen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}
