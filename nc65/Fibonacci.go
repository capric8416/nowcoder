func Fibonacci(n int) int {
	// write code here
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}

	return a
}