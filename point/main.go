package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x, y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	PrintInt(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	PrintInt(points.y)
	z01.PrintRune('\n')
}

func PrintInt(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	PrintUInt(n)
}

func PrintUInt(n int) {
	if n >= 10 {
		PrintUInt(n / 10)
		n = n % 10
	}
	z01.PrintRune(rune('0' + n%10))
}
