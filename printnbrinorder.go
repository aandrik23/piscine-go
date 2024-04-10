package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0{
		z01.PrintRune('0')
		return 
	}
	var digits [10]int
    for n > 0 {
        digit := n % 10
        digits[digit]++
        n /= 10
    }

    for i := 0; i < 10; i++ {
        for j := 0; j < digits[i]; j++ {
            z01.PrintRune(rune('0' + i))
        }
    }
}




































// import (
// 	"fmt"
// 	"sort"
// )

// func PrintNbrInOrder(n int) {
// 	if n == 0 {
// 		fmt.Print("0")
// 		return
// 	}
// 	var digits []int
// 	for n > 0 {
// 		digits = append(digits, n%10)
// 		n = n / 10
// 		sort.Ints(digits)
// 		for _, digit := range digits {
// 			fmt.Print(digit)		
// 		}
// 	}
// }
