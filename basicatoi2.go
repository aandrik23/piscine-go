package piscine

func BasicAtoi2(s string) int {
	a := 0
	for _, i := range s {
		b := 0
		if i < '0' || i > '9' {
			return 0
		}
		for k := '1'; k <= i; k++ {
			b++
		}
		a = a*10 + b
	}
	return a
}
