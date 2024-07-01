package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	comb := make([]int, n)
	for i := 0; i < n; i++ {
		comb[i] = i
	}

	printCombination(comb)
	for nextCombination(comb, n) {
		printCombination(comb)
	}
}

func printCombination(comb []int) {
	for _, num := range comb {
		z01.PrintRune(rune(num + '0'))
	}

	if comb[0] < 10-len(comb) {
		z01.PrintRune(',')
		z01.PrintRune(' ')
	} else {
		z01.PrintRune('\n')
	}
}

func nextCombination(comb []int, n int) bool {
	for i := n - 1; i >= 0; i-- {
		if comb[i] != 9-(n-1-i) {
			comb[i]++
			for j := i + 1; j < n; j++ {
				comb[j] = comb[j-1] + 1
			}
			return true
		}
	}
	return false
}