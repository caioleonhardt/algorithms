package algorithms

import "math"

// IsPrime returns whether n is prime or not
func IsPrime(n int64) bool {
	return isPrime(n)
}

func isPrime(n int64) bool {
	if n < 2 {
		return false
	}

	if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	sqr := math.Sqrt(float64(n))

	for i := int64(3); i <= int64(sqr); i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
