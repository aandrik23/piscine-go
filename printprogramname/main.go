package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	i := []rune(os.Args[0])
	for _, j := range i {
		z01.PrintRune(j)
	}
	z01.PrintRune('\n')
}
