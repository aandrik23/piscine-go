package piscine

func StringToIntSlice(str string) []int {
	intSlice := make([]int, 0)

	for _, char := range str {
		intSlice = append(intSlice, int(char))
	}

	return intSlice
}
