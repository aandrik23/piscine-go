package piscine

func IsPrintable(s string) bool {
	for _, i := range []rune(s){
		if (i<'a' || i > 'z') && (i < 'A' || i > 'Z'){
			return false
		}
	}
	return true
}
