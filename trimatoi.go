package piscine

func TrimAtoi(s string) int {
	i := 0
	j := 1
	for _, value := range s {
		if value >= '0' && value <= '9' {
			value = value - '0'
			i = i*10 + int(value)
		} else if i == 0 && value == '-' {
			j = -1
		}
	}
	return j * i
}
