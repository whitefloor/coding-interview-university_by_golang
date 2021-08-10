package main

func main() {
}

// ReverseStringDoNotUseFor 使用遞迴,且不使用for反轉字串
func ReverseStringDoNotUseFor(str string) string {
	strRune := []rune(str)
	if len(strRune) <= 1 {
		return string(strRune)
	}

	return ReverseStringDoNotUseFor(string(strRune[1:])) + string(strRune[0])
}
