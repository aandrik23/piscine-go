package piscine

func Map(f func(int) bool, a []int) []bool {

	lens := 0
	for range a {
		lens++
	}

	index := 0
	array := make([]bool, lens)
	for _, v := range a {
		array[index] = f(v)
		index++
	}
	return array
}
