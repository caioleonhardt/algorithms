package algorithms

// GCD calcules the greatest common divisor
func GCD(a, b int) int {
	return euclidean(a, b)
}

func euclidean(a, b int) int {
	if b == 0 {
		return a
	}

	return euclidean(b, a%b)
}
