package piscine

func Join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	totalLength := 0
	for _, str := range strs {
		totalLength += len(str) + len(sep)
	}
	totalLength -= len(sep)
	result := make([]byte, totalLength)
	offset := 0
	for _, str := range strs {
		copy(result[offset:], str)
		offset += len(str)
		copy(result[offset:], sep)
		offset += len(sep)
	}
	return string(result)
}
