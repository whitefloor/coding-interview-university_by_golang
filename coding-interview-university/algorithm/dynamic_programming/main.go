package main

func main() {

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
