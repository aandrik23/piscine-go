// package piscine

// import "github.com/01-edu/z01"

// func PrintNbr(n int) {
// 	// Handle the edge case for the minimum 64-bit signed integer
// 	if n == -9223372036854775808 {
// 		str := "-9223372036854775808"
// 		for _, ch := range str {
// 			z01.PrintRune(ch)
// 		}
// 		return
// 	}

// 	if n < 0 {
// 		z01.PrintRune('-')
// 		n = -n
// 	}

// 	if n == 0 {
// 		z01.PrintRune('0')
// 		return
// 	}

// 	var digits []rune
// 	for n > 0 {
// 		digit := n % 10
// 		digits = append([]rune{rune('0' + digit)}, digits...)
// 		n /= 10
// 	}

// 	for _, digit := range digits {
// 		z01.PrintRune(digit)
// 	}
// }

package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	if n == 0 {
		z01.PrintRune('0')
		return
	}

	var digits []rune
	for n > 0 {
		digit := n % 10
		digits = append([]rune{rune('0' + digit)}, digits...)
		n /= 10
	}

	for _, digit := range digits {
		z01.PrintRune(digit)
	}
}
