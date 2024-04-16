package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	piscine.ForEach(PrintNbr, a)
}

func PrintNbr(n int) {
	fmt.Print(n)
}
