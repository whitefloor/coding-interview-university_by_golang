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

// Fibonacci 算出第N個數列的值
func Fibonacci(n int) int {
	x1, y2, z3 := 0, 1, 0

	if n == 0 {
		return x1
	}
	if n == 1 {
		return y2
	}

	for i := 1; i < n; i++ {
		z3 = x1 + y2
		x1 = y2
		y2 = z3
	}

	return z3
}
