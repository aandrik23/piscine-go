package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for a := 0; a <= 98; a++ {
		for b := a + 1; b <= 99; b++ {
			printNumber(a)
			z01.PrintRune(' ')
			printNumber(b)
			if a != 98 || b != 99 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}

	z01.PrintRune('\n')
}

func printNumber(num int) {
	z01.PrintRune(rune((num/10)%10 + '0'))
	z01.PrintRune(rune(num%10 + '0'))
}
