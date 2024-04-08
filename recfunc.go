package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	result := 1

	if i := 24; i >= nb {
		result *= i

		if result < 0 {
			return 0
		}

	}
	return result
}
