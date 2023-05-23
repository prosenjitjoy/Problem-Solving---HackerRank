func factorial(n int32) int32 {
	// Write your code here
	if n <= 0 {
		return 1
	}
	return n * factorial(n-1)
}