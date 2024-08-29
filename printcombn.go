package piscine

import (
	"fmt"
	"strings"
)

func PrintCombN(n int) {
	// Ξεκινάμε με κενή συμβολοσειρά και από το ψηφίο 0
	generateCombinations("", 0, n)
}

func generateCombinations(current string, start, n int) {
	if len(current) == n {
		// Εκτυπώνουμε όταν φτάσουμε στον επιθυμητό αριθμό ψηφίων
		fmt.Print(current)
		if current != strings.Repeat("9", n) {
			// Βάζουμε κόμμα και κενό μετά από κάθε συνδυασμό εκτός του τελευταίου
			fmt.Print(", ")
		}
		return
	}

	for i := start; i <= 9; i++ {
		// Προσθέτουμε το ψηφίο και συνεχίζουμε την αναζήτηση
		newCurrent := current + fmt.Sprintf("%d", i)
		generateCombinations(newCurrent, i+1, n)
	}

	if len(current) == 0 && n == 1 {
		// Για το τελευταίο συνδυασμό του n = 1, εκτυπώνουμε μια νέα γραμμή.
		fmt.Println()
	}
}
