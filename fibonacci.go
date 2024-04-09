package piscine

func Fibonacci(index int) int {
	for index < 0 {
		return -1
	}
	for index == 0 {
		return 0
	}
	for index == 1 {
		return 1
	}
	return Fibonacci(index-2) + Fibonacci(index-1)
}
