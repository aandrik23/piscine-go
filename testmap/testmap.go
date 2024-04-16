package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := piscine.Map(IsPrime, a)
	fmt.Println(result)
}

func IsPrime(nbr int) bool {
	if nbr <= 1 {
		return false
	}
	for i := 2; i*i <= nbr; i++ {
		if nbr%i == 0 {
			return false
		}
	}
	return true
}
