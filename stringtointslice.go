package piscine

func StringToIntSlice(str string) []int {
	if str == "" {
		return nil
	}
	
	intSlice := make([]int, 0)

	for _, char := range str {
		intSlice = append(intSlice, int(char))
	}

	return intSlice
}
