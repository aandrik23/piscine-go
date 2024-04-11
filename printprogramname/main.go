package main

import (
	"os"

	"github.com/01-edu/z01"
)

// test comment
func main() {
	appname := os.Args[0][2:]

	for _, char := range appname {
		z01.PrintRune(rune(char))
	}
	z01.PrintRune('\n')
}
