package piscine

func ConcatParams(args []string) string {
	var i string
	count := 0
	for range args {
		count++
	}
	for j, v := range args {
		if j == count-1 {
			i += v
		} else {
			i += v
			i += "\n"
		}
	}
	return i
}
