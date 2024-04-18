package piscine

func Compact(ptr *[]string) int {
	s := *ptr

	count := 0
	for _, v := range s {
		if v != "" {
			count++
		}
	}

	j := 0
	for i := 0; i < len(s); i++ {
		if s[i] != "" {
			s[j] = s[i]
			j++
		}
	}

	*ptr = s[:count]

	return count
}
