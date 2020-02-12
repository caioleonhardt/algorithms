package algorithms

// Fib calculates the fibonacci number
func Fib(num int) int {
	return memoization(num)
}

func memoization(n int) int {
	if n <= 1 {
		return n
	}

	f := make([]int, n+1)
	f[0] = 0
	f[1] = 1

	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}

	return f[n]
}

func recursion(n int) int {
	if n <= 1 {
		return n
	}

	return recursion(n-1) + recursion(n-2)
}
